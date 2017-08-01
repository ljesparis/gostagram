package gostagram

import "testing"

func TestClient_GetMediaLikes(t *testing.T) {
	client := CreateClient(t)
	users, err := client.GetMediaLikes("1499806890266583125_2451237325")

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

func TestClient_PostMediaLike(t *testing.T) {
	client := CreateClient(t)
	err := client.PostMediaLike("1499806890266583125_2451237325")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("media photo liked.")
}

func TestClient_DeleteMediaLike(t *testing.T) {
	client := CreateClient(t)
	err := client.DeleteMediaLike("1499806890266583125_2451237325")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("media photo disliked.")
}