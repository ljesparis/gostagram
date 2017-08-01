package gostagram

import "testing"

func TestClient_GetCurrentUserRelationship(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetCurrentUserRelationship("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_UnFollowUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.UnFollowUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_FollowUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.FollowUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_BlockUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.BlockUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_UnBlockUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.UnBlockUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_IgnoreUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.IgnoreUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}

func TestClient_ApproveUserById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.ApproveUserById("3066964584")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(DefaultStringValueIfEmpty(tmp.IncomingStatus))
	t.Log(DefaultStringValueIfEmpty(tmp.OutgoingStatus))
}
