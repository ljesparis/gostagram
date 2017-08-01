package gostagram

import (
	"testing"
)

func TestClient_GetMediaComments(t *testing.T) {
	client := CreateClient(t)

	comments, err := client.GetMediaComments("1499806890266583125_2451237325")

	if err != nil {
		t.Fatal(err)
	}

	if len(comments) > 0 {
		for _, comment := range comments {
			t.Log("------------------ Start Comment ------------------")
			t.Log("Id: ", DefaultStringValueIfEmpty(comment.Id))
			t.Log("Text: ", DefaultStringValueIfEmpty(comment.Text))
			t.Log("Created Time: ", DefaultStringValueIfEmpty(comment.CreatedTime))
			t.Log("------------------ Start User Comment ------------------")
			LogUser(&comment.From, t)
			t.Log("------------------ End User Comment ------------------")
			t.Log("------------------ End Comment ------------------")
		}
	} else {
		t.Log("no comments.")
	}
}

func TestClient_PostMediaComment(t *testing.T) {
	client := CreateClient(t)
	err := client.PostMediaComment("api text", "1499806890266583125_2451237325")
	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Comment posted.!")
	}
}

func TestClient_DeleteMediaComment(t *testing.T) {
	client := CreateClient(t)

	err := client.DeleteMediaComment("1499806890266583125_2451237325", "17876627794082799")

	if err != nil {
		t.Fatal(err)
	} else {
		t.Log("Comment deleted.")
	}
}
