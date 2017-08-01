package gostagram

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

type Relationship struct {
	OutgoingStatus string `mapstructure:"outgoing_status"`
	IncomingStatus string `mapstructure:"incoming_status"`
}

func (r Relationship) Follows() bool {
	return r.OutgoingStatus == "follows"
}

func (r Relationship) Requested() bool {
	return r.OutgoingStatus == "requested"
}

func (r Relationship) FollowedBy() bool {
	return r.IncomingStatus == "followed_by"
}

func (r Relationship) RequestedBy() bool {
	return r.IncomingStatus == "requested_by"
}

func (r Relationship) BlockedByYou() bool {
	return r.IncomingStatus == "blocked_by_you"
}


func (c *Client) doRelationshipAction(action, uri string) (*Relationship, error) {
	tmp, err := c.post(uri, BodyData{
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

func (c *Client) GetCurrentUserRelationship(user_id string) (*Relationship, error) {
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

func (c *Client) FollowUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("follow", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) UnFollowUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("unfollow", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) ApproveUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("approve", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) IgnoreUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("ignore", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) BlockUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("block", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

func (c *Client) UnBlockUserById(id string) (*Relationship, error) {
	return c.doRelationshipAction("unblock", fmt.Sprintf("%susers/%s/relationship?access_token=%s", apiUrl, id, c.access_token))
}

