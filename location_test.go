package gostagram

import "testing"

var (
	location_id        = ""
	latitude           = ""
	longitude          = ""
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
	client := CreateClient(t)
	tmp, err := client.SearchLocations(latitude, longitude, nil)
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

func TestClient_SearchLocations_DistanceError2(t *testing.T) {
	FatalIfEmptyString(latitude, "latitude cannot be empty.", t)
	FatalIfEmptyString(longitude, "longitude cannot be empty.", t)
	client := CreateClient(t)
	_, err := client.SearchLocations(latitude, longitude, Parameters{
		"distance": "asd",
	})

	if err != nil {
		t.Log("Success distance most be a digit")
	} else {
		t.Fatal("Error, distance parameter it is not a digit and request was sended anyway.")
	}
}

func TestClient_SearchLocations_DistanceError(t *testing.T) {
	FatalIfEmptyString(latitude, "latitude cannot be empty.", t)
	FatalIfEmptyString(longitude, "longitude cannot be empty.", t)
	client := CreateClient(t)
	_, err := client.SearchLocations(latitude, longitude, Parameters{
		"distance": "800",
	})

	if err != nil {
		t.Log("Success distance cannot be higher to 750")
	} else {
		t.Fatal("Error, distance parameter it is higher to 750 and request was sended anyway.")
	}
}
