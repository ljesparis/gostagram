package gostagram

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Tag struct {
	Name       string
	MediaCount int `mapstructure:"media_count"`
}

func (c *Client) GetTagByName(tagname string) (*Tag, error) {
	tmp, _, err := c.get(fmt.Sprintf("%stags/%s?access_token=%s", apiUrl, tagname, c.access_token))

	if err != nil {
		return nil, err
	}

	tmpTag := (*tmp).(map[string]interface{})
	var tag Tag
	if err = mapstructure.Decode(tmpTag, &tag); err != nil {
		return nil, err
	}

	return &tag, nil
}

func (c *Client) SearchTags(query string) ([]*Tag, error) {
	tmp, _, err := c.get(fmt.Sprintf("%stags/search?q=%s&access_token=%s", apiUrl, query, c.access_token))
	if err != nil {
		return nil, err
	}

	tmpUsers := (*tmp).([]interface{})

	var tags []*Tag
	for _, tmpUser := range tmpUsers {
		var tag Tag

		if err := mapstructure.Decode(tmpUser, &tag); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}
