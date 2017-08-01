package gostagram

import (
	"fmt"
	"errors"
	"strings"
	"regexp"

	"github.com/mitchellh/mapstructure"
)

var (
	CommentsUrlExeed            = errors.New("Cannot contain more than 1 URL.")
	CommentsHashtagExceed       = errors.New("Comment cannot contain more than 4 hashtags.")
	CommentsMaxLengthExceed     = errors.New("Comment cannot exceed 300 characters.")
	CommentsCapitalLettersError = errors.New("Comment cannot consist of all capital letters")
)

var (
	// take this pattern it from
	// https://stackoverflow.com/questions/6883049/regex-to-find-urls-in-string-in-python
	// to match an url.
	urlMatcher           = regexp.MustCompile(`http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+`)
	capitalLetterMatcher = regexp.MustCompile(`[A-Z]`)
)

type Comment struct {
	From        User

	Id          string
	Text        string
	CreatedTime string `mapstructure:"created_time"`
}

func (c *Client) GetMediaComments(media_id string) ([]*Comment, error) {
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

func (c *Client) PostMediaComment(text, media_id string) error {
	if len(text) > 300 {
		return CommentsMaxLengthExceed
	} else if strings.Count(text, "#") > 4 {
		return CommentsHashtagExceed
	} else if len(urlMatcher.FindAllSubmatch([]byte(text), -1)) > 1 {
		return CommentsUrlExeed
	} else {
		for _, c := range text {
			if capitalLetterMatcher.Match([]byte(string(c))) {
				return CommentsCapitalLettersError
			}
		}
	}

	_, err := c.post(fmt.Sprintf("%smedia/%s/comments?access_token=%s", apiUrl, media_id, c.access_token), BodyData{
		"text": text,
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteMediaComment(media_id, comment_id string) error {
	_, err := c.delete(fmt.Sprintf("%smedia/%s/comments/%s?access_token=%s", apiUrl, media_id, comment_id, c.access_token))
	if err != nil {
		return err
	}
	return nil
}
