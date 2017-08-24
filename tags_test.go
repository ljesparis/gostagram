package gostagram

import (
	"testing"
)

var (
	tagname  = ""
	tagquery = ""
)

func TestClient_GetTagByName(t *testing.T) {
	FatalIfEmptyString(tagname, "tagname cannot be empty.", t)

	client := CreateClient(t)
	tmp, err := client.GetTagByName(tagname)
	PanicIfError(err, t)
	t.Log(tmp.Name)
	t.Log(tmp.MediaCount)
}

func TestClient_SearchTags(t *testing.T) {
	FatalIfEmptyString(tagquery, "tagquery cannot be empty.", t)

	client := CreateClient(t)
	tmp, err := client.SearchTags(tagquery)
	PanicIfError(err, t)
	for _, tag := range tmp {
		t.Log("------------------ Start Tag ------------------")
		t.Log(tag.Name)
		t.Log(tag.MediaCount)
		t.Log("------------------ End Tag ------------------")
	}
}
