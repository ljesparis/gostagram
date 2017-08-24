package gostagram

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

// Relationship struct represent
// relationship status with another user.
type Relationship struct {
	OutgoingStatus string `mapstructure:"outgoing_status"`
	IncomingStatus string `mapstructure:"incoming_status"`
}

// check if current user follows other
// user
func (r Relationship) Follows() bool {
	return r.OutgoingStatus == "follows"
}

// check if current user send a request
// to follow other user.
func (r Relationship) Requested() bool {
	return r.OutgoingStatus == "requested"
}

// check if current user was followed by
// other user
func (r Relationship) FollowedBy() bool {
	return r.IncomingStatus == "followed_by"
}

// check if current user was requested by
// other user to accept the follow.
func (r Relationship) RequestedBy() bool {
	return r.IncomingStatus == "requested_by"
}

// check if current user was blocked by
// other user.
func (r Relationship) BlockedByYou() bool {
	return r.IncomingStatus == "blocked_by_you"
}

// Send a post request, with an especific action
// like follow, unfollow, ignore, block, unblock and approve,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/relationships/#post_relationship
func (c Client) doRelationshipAction(action, uri string) (*Relationship, error) {
	tmp, err := c.post(uri, bodyData{
		"action": action,
	})

	if err != nil {
		return nil, err
	}

	tmpRel := (*tmp).(map[string]interface{})
	var relationship Relationship
	if err = mapstructure.Decode(tmpRel, &relationship); err != nil {
		return nil, err
	}

	return &relationship, nil
}

// Check current user relationship with other user by it's id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/relationships/#get_relationship
func (c Client) GetCurrentUserRelationship(user_id string) (*Relationship, error) {
	tmp, _, err := c.get(fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, user_id, c.access_token))
	if err != nil {
		return nil, err
	}

	tmpRel := (*tmp).(map[string]interface{})
	var relationship Relationship
	if err = mapstructure.Decode(tmpRel, &relationship); err != nil {
		return nil, err
	}

	return &relationship, nil
}

func (c Client) FollowUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("follow", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}

func (c Client) UnFollowUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("unfollow", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}

func (c Client) ApproveUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("approve", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}

func (c Client) IgnoreUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("ignore", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}

func (c Client) BlockUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("block", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}

func (c Client) UnBlockUserById(user_id string) (*Relationship, error) {
	return c.doRelationshipAction("unblock", fmt.Sprintf("%susers/%s/relationship?access_token=%s",
		apiUrl, user_id, c.access_token))
}
