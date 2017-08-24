package gostagram

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// UserDetailed struct represents user profile information.
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

// User struct represents main user information,
// most of the time, returned by comments or media
// metadata.
type User struct {
	Id             string
	Type           string
	Username       string
	LastName       string `mapstructure:"last_name"`
	FirstName      string `mapstructure:"first_name"`
	ProfilePicture string `mapstructure:"profile_picture"`
}

func (c Client) getUser(uri string) (*UserDetailed, error) {
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

func (c Client) getUsers(uri string) ([]*User, error) {
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

// Get current logged user data,
// for more information go to
// https://www.instagram.com/developer/endpoints/users/#get_users_self
func (c Client) GetCurrentUser() (*UserDetailed, error) {
	return c.getUser(fmt.Sprintf("%susers/self/?access_token=%s", apiUrl, c.access_token))
}

// Get other user data,
// got more information go to
// https://www.instagram.com/developer/endpoints/users/#get_users
func (c Client) GetUser(id string) (*UserDetailed, error) {
	return c.getUser(fmt.Sprintf("%susers/%s/?access_token=%s", apiUrl, id, c.access_token))
}

// get a list of users from the query,
// for more information go to
// https://www.instagram.com/developer/endpoints/users/#get_users_search
func (c Client) SearchUsers(query string, parameters Parameters) ([]*User, error) {
	tmp := "%susers/search?q=%s&access_token=%s"
	if parameters != nil {
		if parameters["count"] != "" {
			tmp += fmt.Sprintf("&count=%s", parameters["count"])
		}
	}

	return c.getUsers(fmt.Sprintf(tmp,
		apiUrl, query, c.access_token))
}

// Get a list of users that current
// user follows, for more information go to
// https://www.instagram.com/developer/endpoints/relationships/#get_users_follows
func (c Client) GetCurrentUserFollows() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/follows?access_token=%s",
		apiUrl, c.access_token))
}

// Get a list of users, that current user
// is followed, for more information go to
// https://www.instagram.com/developer/endpoints/relationships/#get_users_followed_by
func (c Client) GetCurrentUserFollowedBy() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/followed-by?access_token=%s",
		apiUrl, c.access_token))
}

// Get a list of users who have requested
// this user's permission to follow,
// for more information go to
// https://www.instagram.com/developer/endpoints/relationships/#get_incoming_requests
func (c Client) GetCurrentUserRequestedBy() ([]*User, error) {
	return c.getUsers(fmt.Sprintf("%susers/self/requested-by?access_token=%s",
		apiUrl, c.access_token))
}
