package gostagram

import (
	"fmt"
	"errors"
	"encoding/json"
	url2 "net/url"
	"strings"

	"github.com/parnurzeal/gorequest"
)

const (
	apiUrl    = "https://api.instagram.com/v1/"
	version   = "1.0.0-alpha2"
	userAgent = "gostagram/v1"
)

type Params   map[string]string
type BodyData map[string]interface{}
type Response map[string]interface{}

type Client struct {
	clientSecret     string
	access_token     string

	sandboxMode      bool
	signedRequest    bool
}

func NewClient(access_token string) *Client {
	return &Client{
		clientSecret:     "",
		access_token:     access_token,
		signedRequest:    false,
		sandboxMode:      false,
	}
}

func (c *Client) Version() string {
	return version
}

func (c *Client) SetClientSecret(cs string) {
	c.clientSecret = cs
}

func (c *Client) SetSignedRequest(sr bool) {
	c.signedRequest = sr
}

func (c *Client) SetSandboxMode(sm bool) {
	c.sandboxMode = sm
}

func (c *Client) newRequest(method, uri string, response *Response, dataToSend ...BodyData) error {
	if c.signedRequest && len(c.clientSecret) == 0 {
		return errors.New("Client secret not set.")
	} else if c.signedRequest {

		tmpUrl, err := url2.Parse(uri)

		if err != nil {
			return err
		}

		endpoint := strings.Replace(tmpUrl.EscapedPath(), "/v1", "", 1)

		if endpoint[len(endpoint) - 1] == '/' {
			endpoint = endpoint[0:len(endpoint) - 1]
		}

		params := make(Params)
		for paramName, paramValue := range tmpUrl.Query() {
			params[paramName] = paramValue[0]
		}

		sig, err := c.generateSignature(endpoint, params)

		if err != nil {
			return err
		}

		uri = uri + "&sig=" + sig
	}

	request := gorequest.New().Set("User-Agent", userAgent)

	switch method {
	case gorequest.POST:
		request = request.Post(uri).Type("multipart")

		if dataToSend != nil && len(dataToSend) > 0 {
			request = request.SendMap(dataToSend[0])
		}

		break
	case gorequest.GET:
		request = request.Get(uri).Type("json")
		break
	case gorequest.DELETE:
		request = request.Delete(uri).Type("json")
		break
	}

	_, body, errs := request.EndBytes()
	json.Unmarshal(body, &response)

	if len(errs) > 0 {
		return errs[0]
	}

	if len(*response) > 0 {
		var code int
		var errorType, errorMessage string

		if  (*response)["meta"] != nil {
			meta := (*response)["meta"].(map[string]interface{})
			code = int(meta["code"].(float64))

			if code != 200 {
				errorType = meta["error_type"].(string)
				errorMessage = meta["error_message"].(string)
			} else {
				errorType = ""
				errorMessage = ""
			}
		} else {
			code = int((*response)["code"].(float64))
			errorType = (*response)["error_type"].(string)
			errorMessage = (*response)["error_message"].(string)
		}

		if code != 200 {
			return errors.New(fmt.Sprintf("[%s]: %s", errorType, errorMessage))
		}
	} else {
		return errors.New("Wrong endpoint.")
	}

	return nil
}

func (c *Client) get(url string) (*interface{}, *interface{}, error) {
	var response Response
	err := c.newRequest(gorequest.GET, url, &response)

	if err != nil {
		return nil, nil, err
	}

	tmp := response["data"]
	tmpPagination := response["pagination"]
	if tmpPagination != nil {
		return &tmp, nil, nil
	}

	return &tmp, &tmpPagination, nil
}

func (c *Client) post(url string, dataToSend BodyData) (*interface{}, error) {
	var response Response
	err := c.newRequest(gorequest.POST, url, &response, dataToSend)

	if err != nil {
		return nil, err
	}

	tmp := response["data"]
	return &tmp, nil
}

func (c *Client) delete(url string) (*interface{}, error) {
	var response Response
	err := c.newRequest(gorequest.DELETE, url, &response)

	if err != nil {
		return nil, err
	}

	tmp := response["data"]
	return &tmp, nil
}
