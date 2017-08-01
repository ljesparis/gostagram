package gostagram

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type Location struct {
	Id        string
	Name      string
	Latitude  float64
	Longitude float64
}

func (c *Client) GetLocationById(location_id string) (*Location, error) {
	tmp, _, err := c.get(fmt.Sprintf("%slocations/%s?access_token=%s", apiUrl, location_id, c.access_token))
	if err != nil {
		return nil, err
	}

	tmpLocation := (*tmp).(map[string]interface{})
	var location Location
	if err = mapstructure.Decode(tmpLocation, &location); err != nil {
		return nil, err
	}

	return &location, nil
}

func GetLocationOfRecentMedia(max_id, min_id int, access_token string)  {

}

func (c *Client) SearchLocations(latitude, longitude, distance, facebook_places_id string) ([]*Location, error) {
	tmp, _, err := c.get(fmt.Sprintf("%slocations/search?lat=%s&lng=%s&distance=%s&facebook_places_id=%s&access_token=%s",
		apiUrl, latitude, longitude, distance, facebook_places_id, c.access_token))

	if err != nil {
		return nil, err
	}

	tmpLocations := (*tmp).([]interface{})

	var locations []*Location
	for _, tmplocation := range tmpLocations {
		var location Location

		if err := mapstructure.Decode(tmplocation, &location); err != nil {
			return nil, err
		}

		locations = append(locations, &location)
	}

	return locations, nil
}
