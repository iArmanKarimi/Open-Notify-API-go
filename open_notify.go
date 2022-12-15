package open_notify

import (
	"time"
)

var URLS = struct{ ASTROS, ISS_NOW string }{
	"http://api.open-notify.org/astros.json",
	"http://api.open-notify.org/iss-now.json",
}

type ISSLocation struct {
	message   string    `json:"message"`
	latitude  int       `json:"latitude"`
	longitude int       `json:"longitude"`
	date_time time.Time `json:"timestamp"`
}
type PeopleInSpace struct {
	number  uint   `json:"number"`
	message string `json:"message"`
	people  []struct {
		name  string
		craft string
	} `json:"people"`
}

func GetISSLocation() (*ISSLocation, error) {
	return nil, nil
}

func GetPeoplePeopleInSpace() (*PeopleInSpace, error) {
	return nil, nil
}
