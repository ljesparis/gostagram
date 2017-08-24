package gostagram

import "testing"

func LogRelationShip(rel *Relationship, t *testing.T) {
	t.Log(DefaultStringValueIfEmpty(rel.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(rel.OutgoingStatus))
}

func TestClient_GetCurrentUserRelationship(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.GetCurrentUserRelationship(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_UnFollowUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.UnFollowUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_FollowUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.FollowUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_BlockUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.BlockUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_UnBlockUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.UnBlockUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_IgnoreUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.IgnoreUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}

func TestClient_ApproveUserById(t *testing.T) {
	FatalIfEmptyString(other_user_id, "user id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.ApproveUserById(other_user_id)
	PanicIfError(err, t)
	LogRelationShip(tmp, t)
}
