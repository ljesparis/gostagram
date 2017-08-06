package gostagram

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/mitchellh/mapstructure"
)

const (
	maxUrlPerComment     = 1
	maxCommentLength     = 300
	maxHashTagPerComment = 4
)

var (
	// errors that should be throw
	// if you want to post a comment
	// incorrectly.
	CommentsUrlExceed           = errors.New("Comment cannot contain more than 1 URL.")
	CommentsHashtagExceed       = errors.New("Comment cannot contain more than 4 hashtags.")
	CommentsMaxLengthExceed     = errors.New("Comment cannot exceed 300 characters.")
	CommentsCapitalLettersError = errors.New("Comment cannot consist of all capital letters")
	MissingCommentError         = errors.New("Comment cannot be empty.")
)

var (
	// this url expression was take it from
	// https://stackoverflow.com/questions/6883049/regex-to-find-urls-in-string-in-python.
	urlMatcher = regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
)

// Media resource comment representation.
type Comment struct {
	From        User
	Id          string
	Text        string
	CreatedTime string `mapstructure:"created_time"`
}

// Get all comments from a media resource,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/comments/#get_media_comments
func (c Client) GetMediaComments(media_id string) ([]*Comment, error) {
	tmp, _, err := c.get(fmt.Sprintf("%smedia/%s/comments?access_token=%s", apiUrl, media_id, c.access_token))
	if err != nil {
		return nil, err
	}

	tmpComments := (*tmp).([]interface{})
	var comments []*Comment
	for _, tmpComment := range tmpComments {
		var comment Comment

		if err := mapstructure.Decode(tmpComment, &comment); err != nil {
			return nil, err
		}

		comments = append(comments, &comment)
	}

	return comments, nil
}

// To post a comment in a media resource,
// we need to follow some instagram rules,
// you can find them at https://www.instagram.com/developer/endpoints/comments/#post_media_comments.
func (c Client) PostMediaComment(text, media_id string) error {
	if len(text) > maxCommentLength {
		return CommentsMaxLengthExceed
	} else if strings.Count(text, "#") > maxHashTagPerComment {
		return CommentsHashtagExceed
	} else if len(urlMatcher.FindAllSubmatch([]byte(text), -1)) > maxUrlPerComment {
		return CommentsUrlExceed
	} else if len(text) == 0 {
		return MissingCommentError
	} else {
		for _, c := range text {
			// character most be lowercase.
			if c >= 65 && c <= 90 {
				return CommentsCapitalLettersError
			}
		}
	}

	_, err := c.post(fmt.Sprintf("%smedia/%s/comments?access_token=%s", apiUrl, media_id, c.access_token),
		bodyData{
			"text": text,
		},
	)

	if err != nil {
		return err
	}

	return nil
}

// Delete a comment from a media resource
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/comments/#delete_media_comments
func (c Client) DeleteMediaComment(media_id, comment_id string) error {
	_, err := c.delete(fmt.Sprintf("%smedia/%s/comments/%s?access_token=%s", apiUrl, media_id, comment_id, c.access_token))

	if err != nil {
		return err
	}

	return nil
}
