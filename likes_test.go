package gostagram

import "testing"

func TestClient_GetMediaLikes(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	users, err := client.GetMediaLikes(media_id)
	PanicIfError(err, t)
	IterateUsers(users, t)
}

func TestClient_PostMediaLike(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	err := client.PostMediaLike(media_id)
	PanicIfError(err, t)
	t.Log("media photo liked.")
}

func TestClient_DeleteMediaLike(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	err := client.DeleteMediaLike(media_id)
	PanicIfError(err, t)
	t.Log("media photo disliked.")
}
