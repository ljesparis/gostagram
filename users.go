package gostagram

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type UserDetailed struct {
	Id             string
	Bio            string
	Website        string
	Username       string
	FullName       string `mapstructure:"full_name"`
	ProfilePicture string `mapstructure:"profile_picture"`

	Counts struct {
		Media      int
		Follows    int
		FollowedBy int `mapstructure:"followed_by"`
	}
}

type User struct {
	Id             string
	Type           string
	Username       string
	LastName       string `mapstructure:"last_name"`
	FirstName      string `mapstructure:"first_name"`
	ProfilePicture string `mapstructure:"profile_picture"`
}

func (c *Client) getUser(uri string) (*UserDetailed, error) {
	tmp, _, err := c.get(uri)

	if err != nil {
		return nil, err
	}

	userDetailedMap := (*tmp).(map[string]interface{})
	var userDetailed UserDetailed
	if err = mapstructure.Decode(userDetailedMap, &userDetailed); err != nil {
		return nil, err
	}

	return &userDetailed, nil
}

func (c *Client) getUsers(uri string) ([]*User, error) {
	tmp, _, err := c.get(uri)
	if err != nil {
		return nil, err
	}

	tmpUsers := (*tmp).([]interface{})

	var users []*User
	for _, tmpUser := range tmpUsers {
		var user User

		if err := mapstructure.Decode(tmpUser, &user); err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (c *Client) GetCurrentUser() (*UserDetailed, error) {
	return c.getUser(fmt.Sprintf("%susers/self/?access_token=%s", apiUrl, c.access_token))
}

func (c *Client) GetUser(id string) (*UserDetailed, error) {
	return c.getUser(fmt.Sprintf("%susers/%s/?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) SearchUsers(query string, count int) ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/search?q=%s&count=%d&access_token=%s",
		apiUrl, query, count, c.access_token))
}

func (c *Client) GetCurrentUserFollows() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/follows?access_token=%s",
		apiUrl, c.access_token))
}

func (c *Client) GetCurrentUserFollowedBy() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/followed-by?access_token=%s",
		apiUrl, c.access_token))
}

func (c *Client) GetCurrentUserRequestedBy() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/requested-by?access_token=%s",
		apiUrl, c.access_token))
}
