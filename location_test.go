package gostagram

import "testing"

func TestClient_GetLocationById(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.GetLocationById("1898148680457787")

	if err != nil {
		t.Fatal(err)
	}

	t.Log(tmp.Name)
	t.Log(tmp.Id)
	t.Log(tmp.Longitude)
	t.Log(tmp.Latitude)
}

func TestClient_SearchLocations(t *testing.T) {
	client := CreateClient(t)
	tmp, err := client.SearchLocations("12.12", "23.1", "12312", "12312")

	if err != nil {
		t.Fatal(err)
	}

	for _, location := range tmp {
		t.Log("------------------ Start Location ------------------")
		t.Log(location.Name)
		t.Log(location.Id)
		t.Log(location.Longitude)
		t.Log(location.Latitude)
		t.Log("------------------ End Location ------------------")
	}

}
