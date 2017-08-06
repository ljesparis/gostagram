package gostagram

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

// Tag struct represent
// instagram hastag representation.
type Tag struct {
	Name       string
	MediaCount int `mapstructure:"media_count"`
}

// Get tags(name and how many times was use it), by name
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/tags/#get_tags
func (c Client) GetTagByName(tagname string) (*Tag, error) {
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

// Search tags by a query,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/tags/#get_tags_search
func (c Client) SearchTags(query string) ([]*Tag, error) {
	tmp, _, err := c.get(fmt.Sprintf("%stags/search?q=%s&access_token=%s", apiUrl, query, c.access_token))
	if err != nil {
		return nil, err
	}

	tmpTags := (*tmp).([]interface{})

	var tags []*Tag
	for _, tmpTag := range tmpTags {
		var tag Tag

		if err := mapstructure.Decode(tmpTag, &tag); err != nil {
			return nil, err
		}

		tags = append(tags, &tag)
	}

	return tags, nil
}
