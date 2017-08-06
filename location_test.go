package gostagram

import "testing"

var (
	location_id        = ""
	latitude           = ""
	longitude          = ""
	distance           = ""
	facebook_places_id = ""
)

func TestClient_GetLocationById(t *testing.T) {
	FatalIfEmptyString(location_id, "Location id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.GetLocationById(location_id)
	PanicIfError(err, t)
	t.Log(tmp.Name)
	t.Log(tmp.Id)
	t.Log(tmp.Longitude)
	t.Log(tmp.Latitude)
}

func TestClient_SearchLocations(t *testing.T) {
	FatalIfEmptyString(latitude, "latitude cannot be empty.", t)
	FatalIfEmptyString(longitude, "longitude cannot be empty.", t)
	FatalIfEmptyString(distance, "distance cannot be empty.", t)
	FatalIfEmptyString(facebook_places_id, "facebook places id cannot be empty.", t)
	client := CreateClient(t)
	tmp, err := client.SearchLocations(latitude, longitude, distance, facebook_places_id)
	PanicIfError(err, t)
	for _, location := range tmp {
		t.Log("------------------ Start Location ------------------")
		t.Log(location.Name)
		t.Log(location.Id)
		t.Log(location.Longitude)
		t.Log(location.Latitude)
		t.Log("------------------ End Location ------------------")
	}
}
