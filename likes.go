package gostagram

import "fmt"

func (c *Client) GetMediaLikes(media_id string) ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token))
}

func (c *Client) PostMediaLike(media_id string) error {
	_, err := c.post(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token), nil)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteMediaLike(media_id string) error {
	_, err := c.delete(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token))
	if err != nil {
		return err
	}
	return nil
}
