package gostagram

import (
	"fmt"
	"errors"

	"github.com/mitchellh/mapstructure"
	"strconv"
)

var (
	maxLocationDistanceError = errors.New("Maximun distance is 5km.")
)

const (
	maxLocationDistance = 750
)

type Location struct {
	Id        string
	Name      string
	Latitude  float64
	Longitude float64
}

// Get an especific location by a respective id,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/locations/#get_locations
func (c Client) GetLocationById(location_id string) (*Location, error) {
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

// Search locations by its latitude, longitude, distance and facebook places,
// for more information about it, go to
// https://www.instagram.com/developer/endpoints/locations/#get_locations_search
func (c Client) SearchLocations(latitude, longitude string, params Parameters) ([]*Location, error) {
	tmp1 := "%slocations/search?lat=%s&lng=%s&access_token=%s"
	if params != nil {
		if params["distance"] != "" {
			distance, err := strconv.Atoi(params["distance"])
			if err != nil {
				return nil, err
			}

			if distance > maxLocationDistance {
				return nil, maxLocationDistanceError
			}

			tmp1 += fmt.Sprintf("&distance=%d", distance)
		}
	}

	tmp, _, err := c.get(fmt.Sprintf(tmp1,
		apiUrl, latitude, longitude, c.access_token))

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
