package iatacodes

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Airline struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Airport struct {
	Code string `json:"code"`
}

type LegTime struct {
	Scheduled       time.Time `json:"scheduled"`
	LatestPublished time.Time `json:"latestPublished"`
	Estimated       struct {
		Value time.Time `json:"value"`
	} `json:"estimated"`
}

type LegInformation struct {
	Airport Airport `json:"airport"`
	Times   LegTime `json:"times"`
}

type FlightLeg struct {
	StatusName           string         `json:"statusName"`
	PublishedStatus      string         `json:"publishedStatus"`
	DepartureInformation LegInformation `json:"departureInformation"`
	ArrivalInformation   LegInformation `json:"arrivalInformation"`
}

type FlightStatus struct {
	FlightNumber int         `json:"flightNumber"`
	Route        []string    `json:"route"`
	Airline      Airline     `json:"airline"`
	FlightLegs   []FlightLeg `json:"flightLegs"`
}

func (ic *IATACodes) FlightStatusList(start time.Time, end time.Time) ([]FlightStatus, error) {

	fs := []FlightStatus{}

	url := fmt.Sprintf("/v7/routes?limit=10&api-key=%s", ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []FlightStatus{}, err
	}

	log.Printf("GOT: %s\n", resp)

	d := struct {
		FlightStatus []FlightStatus `json:"operationalFlights"`
		Page         struct {
			PageSize   int `json:"pageSize"`
			PageNumber int `json:"pageNumber"`
			FullCount  int `json:"fullCount"`
			PageCount  int `json:"pageCount"`
			TotalPages int `json:"totalPages"`
		} `json:"page"`
	}{}

	if err := json.Unmarshal([]byte(resp), &d); err != nil {
		return []FlightStatus{}, err
	}

	return fs, nil

}
