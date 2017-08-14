package gostagram

import (
	"testing"
)

var (
	comment_id = ""
)

func TestClient_GetMediaComments(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	comments, err := client.GetMediaComments(media_id)
	PanicIfError(err, t)

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

// post a comment.
func TestClient_PostMediaComment(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	err := client.PostMediaComment("api text", media_id)
	PanicIfError(err, t)
	t.Log("Comment posted.!")
}

// catching following errors:
//  - uppercase error.
//  - max length error.
//  - max hashtag per comment.
//  - max url per comment.
//  - missing comment or empty comment.
func TestClient_PostMediaComment2(t *testing.T) {
	var comment string = ""
	for i := 0; i <= 400; i++ {
		comment += "a"
	}

	input := []string {
		comment,
		"Api test",
		"aPi test",
		"apI test",
		"api Test",
		"apI tEst",
		"apI teSt",
		"apI tesT",
		"apI tEst",
		"#apI t#st",
		"you can find a go client for instagram at https://github.com/leoxnidas/gostagram and this is my personal github https://github.com/leoxnidas.",
		"#new #instagram #api #for #go",
		"",
	}

	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	client := CreateClient(t)
	hasErr := true

	for _, el := range input {
		err := client.PostMediaComment(el, media_id)
		if err == CommentsCapitalLettersError {
			t.Logf("Cannot comment '%s'.", el)
			hasErr = false
		} else if err == CommentsMaxLengthExceed {
			t.Log("Comment cannot has more than 300 characters.")
			hasErr = false
		} else if err == CommentsHashtagExceed {
			t.Log("Comment cannot has more than 4 hashtags.")
			hasErr = false
		} else if err == CommentsUrlExceed {
			t.Log("Comment cannot has more than 2 urls")
			hasErr = false
		} else if err == MissingCommentError {
			t.Log("Comment cannot be an empty string.")
			hasErr = false
		} else {
			t.Fatal(err)
		}
	}

	if hasErr {
		t.Fatal("Not error throwed!")
	}
}

func TestClient_DeleteMediaComment(t *testing.T) {
	FatalIfEmptyString(media_id, "media id cannot be empty.", t)
	FatalIfEmptyString(comment_id, "comment id cannot be empty.", t)
	client := CreateClient(t)
	err := client.DeleteMediaComment(media_id, comment_id)
	PanicIfError(err, t)
	t.Log("Comment deleted.")
}
