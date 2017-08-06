// +build

package main

import (
	"net/http"
	"fmt"
	"errors"
	"encoding/json"
	"io"
	"net/url"
	"os"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
	"sync"
	"strings"
	"time"
	"reflect"
	"encoding/base64"
	"context"
	"io/ioutil"
	"path"
	"html/template"

	"github.com/ljesparis/gostagram"
)

// ROUTES.
const (
	homeRoute     = "/"
	error404Route = "/error"
	logoutRoute   = "/logout"
	oauth2Route   = "/instagram/access"
)

// ERRORS.
var (
	codeMissing           = errors.New("Authorization code missing.")
	clientIdMissing       = errors.New("Missing client id.")
	redirectUrlMissing    = errors.New("Missing redirect url.")
	accessTokenMissing    = errors.New("Token Misteriously missing.")
	clientSecretMissing   = errors.New("Missing client secret.")
	TemplateDoesNotExists = errors.New("Template does not exists.")
)

// SETTINGS.
var (
	// base
	port = os.Getenv("PORT")
	host = os.Getenv("HOST")

	// templates dir
	templates_dir = os.Getenv("PWD") + "/templates"

	// static
	static_url = "/static/"
	static_dir = os.Getenv("PWD") + "/static"

	// oauth2 (instagram).
	sig_secret   = "" // can be anything you want.
	clientId     = ""
	clientSecret = ""
	redirectUrl  = "http://" + host + ":" + port + oauth2Route
	scopes       = []string {
		"basic",
		"comments",
		"follower_list",
		"likes",
		"public_content",
		"relationships",
	}

	// cookie
	cookie_name    = "gostagram"
	cookie_maxage  = 24 * 60 * 60 // 24 hours.

	// session
	encryption_key = "" //32bytes key
)


//
//
//               INSTAGRAM CLIENT.
//            ------------------------
//
//
//

var (
	client *gostagram.Client
)

//
//
//                   TEMPLATES.
//            ------------------------
//
//
//

type templateContext map[string]interface{}

func readTemplate(filename string) (string, error) {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(dat), nil
}

func findTemplatename(filename string) (string, error) {
	var tmp []string
	if strings.Contains(filename, "/") {
		tmp = strings.Split(filename, "/")
	} else {
		tmp = []string{ filename }
	}

	var templatep string
	for i, p := range tmp {
		templatep = path.Join(templates_dir, p)
		if i == len(tmp) - 1 {
			if _, err := os.Stat(templatep); err == nil && !os.IsNotExist(err) {
				return templatep, nil
			}
		}
	}

	return "", TemplateDoesNotExists
}

func render(res http.ResponseWriter, tpl string, data interface{}, funcs template.FuncMap) {
	temp, err := findTemplatename(tpl)

	if err != nil {
		panic(err)
	}

	dat, err := readTemplate(temp)

	if err != nil {
		panic(err)
	}

	templateEngine, err := template.New("gostagram").Funcs(funcs).Parse(dat)
	if err != nil {
		panic(err)
	}

	templateEngine.Execute(res, data)
}

//
//
//                     USERS.
//            ------------------------
//
//
//

func GetUserFromContext(ctx context.Context) BaseUser {
	tmp := ctx.Value("user").(BaseUser)
	return tmp
}

type BaseUser interface {
	IsAuthenticated() bool
}

type User struct {
	Id        string
	Image     string
	Fullname  string
	Username  string
}

func (u User) IsAuthenticated() bool {
	return true
}

type AnonymousUser struct {}

func (u AnonymousUser) IsAuthenticated() bool {
	return false
}

//
//
//                     SESSION.
//            --------------------------
//
//
//

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}


func encodeSession(data *map[string]string) string {
	tmp := ""
	i := 0
	for key, val := range (*data) {
		tmp += key + "~" + val
		if i < len(*data) - 1 {
			tmp += "|"
		}
		i++
	}

	dd, err := encrypt([]byte(tmp), []byte(encryption_key))

	if err != nil {
		panic(err)
	}

	return base64.StdEncoding.EncodeToString(dd)
}

func decodeSession(session string) *map[string]string {
	session_encripted, err := base64.StdEncoding.DecodeString(session)
	if err != nil {
		return nil
	}

	tmp, err := decrypt([]byte(session_encripted), []byte(encryption_key))
	if err != nil {
		return nil
	}

	tmpStr := string(tmp)
	user := make(map[string]string)
	for _, tmp2 := range strings.Split(tmpStr, "|") {
		if len(tmp2) > 1 {
			tmp3 := strings.Split(tmp2, "~")
			key, val := tmp3[0], tmp3[1]
			user[key] = val
		}
	}

	return &user
}

//
//
//                     OAUTH2.
//            ------------------------
//
//
//

var (
	signed_state = generateState()
)

// Avoid CSRF attack.
func generateState() string {
	pwd := os.Getenv("PWD")
	shell := os.Getenv("SHELL")
	user := os.Getenv("USER")
	PATH := os.Getenv("PATH")

	tmp := hmac.New(sha256.New, []byte(
		fmt.Sprintf("%s-%s-%s-%s-%s", pwd, shell, user, PATH, sig_secret),
	))

	return hex.EncodeToString(tmp.Sum(nil))
}

func generateOauthUrl(client_id, redirect_url, state string, scopes []string) (string, error) {
	if len(client_id) == 0 {
		return "", clientIdMissing
	}

	if len(redirect_url) == 0 {
		return "", redirectUrlMissing
	}

	tmpUrl := "https://api.instagram.com/oauth/authorize/?client_id=%s&redirect_uri=%s&response_type=code"
	tmpUrl = fmt.Sprintf(tmpUrl, client_id, redirect_url)

	if len(scopes) >= 1 {
		tmpUrl += "&scope="
	}

	for i, scope := range scopes {
		tmpUrl += scope
		if i < len(scopes) - 1 {
			tmpUrl += "+"
		}
	}

	if len(state) >= 1{
		tmpUrl += "&state=" + signed_state
	}

	return tmpUrl, nil
}

func exchangeToken(client_id, client_secret, redirect_url, code string) (string, map[string]interface{}, error) {
	if len(client_id) == 0 {
		return "", nil, clientIdMissing
	}

	if len(client_secret) == 0 {
		return "", nil, clientSecretMissing
	}

	if len(code) == 0 {
		return "", nil, codeMissing
	}

	if len(redirect_url) == 0 {
		return "", nil, redirectUrlMissing
	}

	tmpUrl := "https://api.instagram.com/oauth/access_token"
	res, err := http.PostForm(tmpUrl, url.Values{
		"client_id": {client_id},
		"client_secret": {client_secret},
		"grant_type": {"authorization_code"},
		"redirect_uri": {redirect_url},
		"code": {code},
	})

	if err != nil {
		return "", nil, err
	}

	defer res.Body.Close()

	response := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil && err != io.EOF {
		return "", nil, err
	}

	if response["access_token"] != nil {
		token := response["access_token"].(string)
		return token, response["user"].(map[string]interface{}), nil
	}

	return "", nil, accessTokenMissing
}

//
//
//              WEB APP CONTROLLERS.
//            ------------------------
//
//
//

func home(res http.ResponseWriter, req *http.Request) {
	tmpUrl, err := generateOauthUrl(clientId, redirectUrl, signed_state, scopes)
	if err != nil {
		panic(err)
	} else {
		res.Header().Set("Content-Type", "text/html")
		res.WriteHeader(200)

		tplCtx := make(templateContext)
		tplCtx["Url"] = tmpUrl
		user := GetUserFromContext(req.Context())
		if user.IsAuthenticated() {
			tplCtx["User"] = user.(User)

			tmpMedia, err := client.GetCurrentUserRecentMedia("1", "1", 1)
			if err != nil {
				panic(err)
			}

			var media []gostagram.MediaImage
			for _, tmp := range tmpMedia {
				media = append(media, (*tmp).(gostagram.MediaImage))
			}

			tplCtx["MediaImages"] = media
		} else {
			tplCtx["User"] = user.(AnonymousUser)
		}

		render(res, "home.html", tplCtx, template.FuncMap{
			"getMediaComment": func(media_id string) []gostagram.Comment {
				tmpComments, err := client.GetMediaComments(media_id)
				if err != nil {
					return []gostagram.Comment{}
				}

				var comments []gostagram.Comment
				for _, tmpComment := range tmpComments {
					comments = append(comments, *tmpComment)
				}

				return comments
			},
		})
	}
}

func instagramLogin(res http.ResponseWriter, req *http.Request) {
	code := req.URL.Query().Get("code")
	state := req.URL.Query().Get("state")

	if state != signed_state {
		http.Redirect(res, req, error404Route, 302)
	}

	if len(code) == 0 {
		http.Redirect(res, req, error404Route, 302)
	} else {
		token, tmp, err := exchangeToken(clientId, clientSecret, redirectUrl, code)

		if err != nil {
			http.Redirect(res, req, error404Route, 302)
		} else {
			client = gostagram.NewClient(token)
			user := map[string]string {
				"image": tmp["profile_picture"].(string),
				"username": tmp["username"].(string),
				"id": tmp["id"].(string),
				"fullname": tmp["full_name"].(string),
				"tk": token,
			}

			http.SetCookie(res, &http.Cookie{
				Value: encodeSession(&user),
				Name: cookie_name,
				HttpOnly: true,
				Path: "/",
				Expires: time.Now().Add(time.Hour * 24),
				MaxAge: cookie_maxage,
			})

			http.Redirect(res, req, homeRoute, 302)
		}
	}
}

func logout(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{
		Name: cookie_name,
		HttpOnly: true,
		Path: "/",
		Expires: time.Now(),
		MaxAge: -1,
	})

	client = nil
	http.Redirect(res, req, homeRoute, 302)
}

func errorHandler(res http.ResponseWriter, req *http.Request, err error) {
	render(res, "404.html", templateContext{
		"ErrorMessage": err.Error(),
	}, template.FuncMap{})
}

//
//
//              WEB APP MIDDLEWARES.
//            ------------------------
//
//
//

var (
	once   sync.Once
	logger *log.Logger
)

func Logger(next http.Handler) http.HandlerFunc {

	once.Do(func(){
		logger = log.New(os.Stdout, "", log.LUTC)
	})

	return func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()

		defer func() {
			remoteAddr := strings.Split(req.RemoteAddr, ":")[0]
			_path := req.URL.Path
			proto := req.Proto
			brenchmark := time.Now().Sub(start)

			// getting response code and content length from response structure.
			elm := reflect.Indirect(reflect.ValueOf(res))
			status_code := elm.FieldByName("status").Int()
			contentLength := elm.FieldByName("written").Int()

			logger.Printf("Remote:{%s} - Path:{%s} - Code:{%d} - Protocol:{%s} - ContentLength:{%d} - Brenchmark:{%s}",
				_path, remoteAddr, status_code, proto, contentLength, brenchmark,
			)
		} ()

		next.ServeHTTP(res, req)
	}
}

func Recover(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		defer func(){
			if r := recover(); r != nil {
				var err error
				switch r := r.(type) {
				case error:
					err = r
				default:
					err = fmt.Errorf("%v", r)
				}

				errorHandler(res, req, err)
			}
		}()

		next.ServeHTTP(res, req)
	}
}

// authenticating user if session exists.
// base64 session.
func AuthUser(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var user BaseUser
		cookie, err := req.Cookie(cookie_name)
		if err != nil {
			user = BaseUser(AnonymousUser{})
		} else {
			tmpUser := decodeSession(cookie.Value)
			user = BaseUser(User{
				Fullname: (*tmpUser)["fullname"],
				Image: (*tmpUser)["image"],
				Username: (*tmpUser)["username"],
				Id: (*tmpUser)["id"],
			})
			client = gostagram.NewClient((*tmpUser)["tk"])
		}

		ctx := context.WithValue(req.Context(), "user", user)
		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	}
}

func Secure(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.Header().Set("X-XSS-Protection", 	"1; mode=block")
		res.Header().Set("Surrogate-Control", "no-store")
		res.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		res.Header().Set("Pragma", "no-cache")
		res.Header().Set("Expires", "0")
		res.Header().Set("X-Frame-Options", "DENY")
		res.Header().Set("X-Content-Type-Options", "nosniff")
		res.Header().Set("Referrer-Policy", "strict-origin")
		next.ServeHTTP(res, req)
	}
}

//
//
//             MAIN SERVER FUNCTION.
//            -----------------------
//
//
//

func main() {
	if len(port) == 0 {
		fmt.Println("port not set.")
		os.Exit(1)
	} else if len(host) == 0 {
		fmt.Println("host not set.")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	// adding routes.
	mux.HandleFunc(homeRoute, home)
	mux.HandleFunc(logoutRoute, logout)
	mux.HandleFunc(oauth2Route, instagramLogin)

	// add staticfiles directory.
	mux.Handle(static_url, http.StripPrefix(static_url, http.FileServer(http.Dir(static_dir))))

	http.ListenAndServe(host+ ":" + port,
		Logger(Recover(Secure(AuthUser(mux)))),
	)
}
