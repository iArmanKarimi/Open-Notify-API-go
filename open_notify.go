package open_notify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var URLS = struct{ PEOPLE, ISS_NOW string }{
	"http://api.open-notify.org/astros.json",
	"http://api.open-notify.org/iss-now.json",
}

type __ISSLocation struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
	Location  struct {
		Latitude  string `json:"latitude"`
		Longitude string `json:"longitude"`
	} `json:"iss_position"`
}
type ISSLocation struct {
	Message  string
	DateTime time.Time
	Location struct {
		Latitude  float64
		Longitude float64
	}
}
type PeopleInSpace struct {
	Number  uint   `json:"number"`
	Message string `json:"message"`
	People  []struct {
		Name  string `json:"name"`
		Craft string `json:"craft"`
	} `json:"people"`
}

func get(url string) (data []byte, err error) {
	var resp *http.Response
	resp, err = http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Gets the current location of ISS (International Space Station)
func GetISSLocation() (*ISSLocation, error) {
	data, err := get(URLS.ISS_NOW)
	if err != nil {
		return nil, err
	}
	temp_iss_loc := new(__ISSLocation)
	err = json.Unmarshal(data, &temp_iss_loc)
	if err != nil {
		return nil, err
	}
	lat, err := strconv.ParseFloat(temp_iss_loc.Location.Latitude, 64)
	if err != nil {
		return nil, err
	}
	long, err := strconv.ParseFloat(temp_iss_loc.Location.Longitude, 64)
	if err != nil {
		return nil, err
	}
	return &ISSLocation{
		Message: temp_iss_loc.Message,
		Location: struct {
			Latitude  float64
			Longitude float64
		}{
			Latitude:  lat,
			Longitude: long,
		},
		DateTime: time.Unix(temp_iss_loc.Timestamp, 0),
	}, nil
}

// Get the current number of people in space. It also returns the names and spacecraft those people are on.
func GetPeopleInSpace() (*PeopleInSpace, error) {
	data, err := get(URLS.PEOPLE)
	if err != nil {
		return nil, err
	}
	ppl := new(PeopleInSpace)
	err = json.Unmarshal(data, &ppl)
	if err != nil {
		return nil, err
	}
	return ppl, nil
}
