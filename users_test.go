package gostagram

import (
	"testing"
)

var (
	// basic variables to create
	// an instagram client.
	clientSecret     = ""
	accessToken      = ""
	useSignedRequest = false

	other_user_id = ""
	user_query = ""
)

func CreateClient(t *testing.T) *Client {
	FatalIfEmptyString(clientSecret, "Access token not set.", t)
	client := NewClient(accessToken)

	if useSignedRequest {
		client.SetSignedRequest(true)
		FatalIfEmptyString(clientSecret, "Client secret not set and is need it for signed request.", t)
		client.SetClientSecret(clientSecret)
	}

	return client
}

// throw an error, if an error exists.
func PanicIfError(err error, t *testing.T) {
	if err != nil {
		t.Fatal(err)
	}
}

// throw an error if b string is empty.
func FatalIfEmptyString(b, errMessage string, t *testing.T) {
	if len(b) == 0 {
		t.Fatal(errMessage)
	}
}

// will replace empty string by 'None'
func DefaultStringValueIfEmpty(t string) string {
	if len(t) > 0 {
		return t
	}

	return "None"
}

func LogUserDetailed(user *UserDetailed, t *testing.T) {
	t.Log("------------------ Start User ------------------")
	t.Log("Username: ", DefaultStringValueIfEmpty(user.Username))
	t.Log("Profile Picture: ", DefaultStringValueIfEmpty(user.ProfilePicture))
	t.Log("Fullname: ", DefaultStringValueIfEmpty(user.FullName))
	t.Log("Website: ", DefaultStringValueIfEmpty(user.Website))
	t.Log("Bio: ", DefaultStringValueIfEmpty(user.Bio))
	t.Log("Id: ", DefaultStringValueIfEmpty(user.Id))
	t.Log("Followed by: ", user.Counts.FollowedBy)
	t.Log("Follows: ", user.Counts.Follows)
	t.Log("Media: ", user.Counts.Media)
	t.Log("------------------ End User ------------------")
}

func LogUser(user *User, t *testing.T) {
	t.Log("Username: ", DefaultStringValueIfEmpty(user.Username))
	t.Log("Profile Picture: ", DefaultStringValueIfEmpty(user.ProfilePicture))
	t.Log("First Name: ", DefaultStringValueIfEmpty(user.FirstName))
	t.Log("Last Name: ", DefaultStringValueIfEmpty(user.LastName))
	t.Log("Id: ", DefaultStringValueIfEmpty(user.Id))
	t.Log("Type: ", DefaultStringValueIfEmpty(user.Type))
}

func IterateUsers(users []*User, t *testing.T) {
	if len(users) > 0 {
		for _, user := range users {
			t.Log("------------------ Start User ------------------")
			LogUser(user, t)
			t.Log("------------------ End User ------------------")
		}
	} else {
		t.Log("Not users found.!")
	}
}

func TestClient_GetCurrentUser(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetCurrentUser()
	PanicIfError(err, t)
	LogUserDetailed(tmp, t)
}

func TestClient_GetUser(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.GetUser(other_user_id)
	PanicIfError(err, t)
	LogUserDetailed(tmp, t)
}

func TestClient_SearchUsers(t *testing.T) {
	client := CreateClient(t)
	users, err := client.SearchUsers(user_query, 1)
	PanicIfError(err, t)
	IterateUsers(users, t)
}

func TestClient_GetCurrentUserFollows(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserFollows()
	PanicIfError(err, t)
	IterateUsers(users, t)
}

func TestClient_GetCurrentUserFollowedBy(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserFollowedBy()
	PanicIfError(err, t)
	IterateUsers(users, t)
}

func TestClient_GetCurrentUserRequestedBy(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserRequestedBy()
	PanicIfError(err, t)
	IterateUsers(users, t)
}
