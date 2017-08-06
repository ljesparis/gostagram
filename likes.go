package gostagram

import "fmt"

// Get 'likes' from media resource with the respective id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/likes/#get_media_likes
func (c Client) GetMediaLikes(media_id string) ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token))
}

// Do like a media resource, with a respective id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/likes/#post_likes
func (c Client) PostMediaLike(media_id string) error {
	_, err := c.post(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token))
	if err != nil {
		return err
	}

	return nil
}

// Delete a media resource, with a respective id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/likes/#delete_likes
func (c Client) DeleteMediaLike(media_id string) error {
	_, err := c.delete(fmt.Sprintf("%smedia/%s/likes?access_token=%s", apiUrl, media_id, c.access_token))
	if err != nil {
		return err
	}

	return nil
}
