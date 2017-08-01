package gostagram

import (
	"testing"
	"os"
)

var (
	clientSecret = "94637a21fdab42dca158c6f6bd1fbe19"
	accessToken  = "2451237325.e0f323b.e8ab7baadad945e0856bebcae2032127"
)

func GetAccessTokenFromEnv() string {
	return os.Getenv("ACCESS_TOKEN")
}

func GetClientSecret() string {
	return os.Getenv("CLIENT_SECRET")
}

func IsSecureRequest() bool {
	return os.Getenv("SECURE") == "true"
}

func GetUserId() string {
	return os.Getenv("USER_ID")
}

func CreateClient(t *testing.T) *Client {
	//accessToken := GetAccessTokenFromEnv()
	//if len(accessToken) == 0 {
	//	t.Error("Access token not set.")
	//}

	client := NewClient(accessToken)

	//if IsSecureRequest() {
	//	client.SetSignedRequest(true)
	//	client_secret := GetClientSecret()
	//	if len(client_secret) == 0 {
	//		t.Error("Cannot send secure request without client secret")
	//	} else {
	//		client.SetClientSecret(client_secret)
	//	}
	//}

	client.SetSignedRequest(true)
	client.SetClientSecret(clientSecret)

	return client
}

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

func TestClient_GetCurrentUser(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetCurrentUser()

	if err != nil {
		t.Fatal(err)
	}

	LogUserDetailed(tmp, t)
}

func TestClient_GetUser(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetUser("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	LogUserDetailed(tmp, t)
}

func TestClient_SearchUsers(t *testing.T) {
	client := CreateClient(t)
	users, err := client.SearchUsers("leo", 1)

	if err != nil {
		t.Fatal(err)
	}

	if len(users) > 0 {
		for _, user := range users {
			t.Log("------------------ Start User ------------------")
			LogUser(user, t)
			t.Log("------------------ End User ------------------")
		}
	} else {
		t.Log("Not found any user.!")
	}
}

func TestClient_GetCurrentUserFollows(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserFollows()

	if err != nil {
		t.Fatal(err)
	}

	if len(users) > 0 {
		for _, user := range users {
			t.Log("------------------ Start User ------------------")
			LogUser(user, t)
			t.Log("------------------ End User ------------------")
		}
	} else {
		t.Log("Not found any user.!")
	}
}

func TestClient_GetCurrentUserFollowedBy(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserFollowedBy()

	if err != nil {
		t.Fatal(err)
	}

	if len(users) > 0 {
		for _, user := range users {
			t.Log("------------------ Start User ------------------")
			LogUser(user, t)
			t.Log("------------------ End User ------------------")
		}
	} else {
		t.Log("Nothing found.!")
	}
}

func TestClient_GetCurrentUserRequestedBy(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetCurrentUserRequestedBy()

	if err != nil {
		t.Fatal(err)
	}

	if len(users) > 0 {
		for _, user := range users {
			t.Log("------------------ Start User ------------------")
			LogUser(user, t)
			t.Log("------------------ End User ------------------")
		}
	} else {
		t.Log("Not found any user.!")
	}
}
