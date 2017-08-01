package gostagram

import (
	"testing"
)

func TestClient_GetTagByName(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetTagByName("leoxnidas")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(tmp.Name)
	t.Log(tmp.MediaCount)
}

func TestClient_SearchTags(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.SearchTags("leoxn")

	if err != nil {
		t.Fatal(err)
	}

	for _, tag := range tmp {
		t.Log("------------------ Start Tag ------------------")
		t.Log(tag.Name)
		t.Log(tag.MediaCount)
		t.Log("------------------ End Tag ------------------")
	}
}