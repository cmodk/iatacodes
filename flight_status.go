package iatacodes

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type ICRequestGeo struct {
	Timezone string `json:"timezone"`
}

type ICRequestClient struct {
	Geo ICRequestGeo `json:"geo"`
}

type ICRequest struct {
	Language string          `json:"lang"`
	Client   ICRequestClient `json:"client"`
}

type RouteResponse struct {
	FlightNumber      string `json:"flight_number"`
	AirlineIATA       string `json:"airline_iata"`
	AirlineICAO       string `json:"airline_icao"`
	DepartureICAO     string `json:"departure_icao"`
	DepartureIATA     string `json:"departure_iata"`
	DepartureTerminal string `json:"departure_terminal"`
	DepartureGate     string `json:"departure_gate"`
	ArrivalICAO       string `json:"arrival_icao"`
	ArrivalIATA       string `json:"arrival_iata"`
	ArrivalTerminal   string `json:"arrival_terminal"`
	ArrivalGate       string `json:"arrival_gate"`
	Codeshares        string `json:"code_shares"`
}

type AirportTable struct {
	IATA          string   `json:"iata_code"`
	ICAO          string   `json:"icao_code"`
	Termial       string   `json:"terminal"`
	Gate          string   `json:"gate"`
	Baggage       string   `json:"baggage"`
	ScheduledTime IATATime `json:"scheduled_time"`
	EstimatedTime IATATime `json:"estimated_time"`
	ActualTime    IATATime `json:"actual_time"`
}

type TimetableResponse struct {
	Type   string `type"`
	Status string `json:"status"`
	Flight struct {
		Number string `json:"number"`
		IATA   string `json:"iata_number"`
		ICAO   string `json:"icao_number"`
	} `json:"flight"`
	Airline struct {
		Name string `json:"name"`
		ICAO string `json:"icao_code"`
		IATA string `json:"iata_code"`
	} `json:"airline"`
	Departure  AirportTable `json:"departure"`
	Arrival    AirportTable `json:"arrival"`
	Codeshares string       `json:"codeshares"`
}

type FlightStatus struct {
	Request  ICRequest       `json:"request"`
	Response []RouteResponse `json:"response"`
}

func (ic *IATACodes) RouteList() ([]RouteResponse, error) {

	fs := FlightStatus{}

	url := fmt.Sprintf("/v7/routes?api_key=%s", ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []RouteResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &fs); err != nil {
		return []RouteResponse{}, err
	}

	return fs.Response, nil

}

type TimetableRequest struct {
	Request  ICRequest           `json:"request"`
	Response []TimetableResponse `json:"response"`
}

func (tr *TimetableRequest) CorrectTimes() {
	if debug {
		log.Printf("Request timezone: %s\n", tr.Request.Client.Geo.Timezone)
	}

	location, err := time.LoadLocation(tr.Request.Client.Geo.Timezone)
	if err != nil {
		panic(err)
	}

	for i, _ := range tr.Response {
		r := &(tr.Response[i])

		r.Departure.ScheduledTime.ChangeTimezone(location)
		r.Departure.EstimatedTime.ChangeTimezone(location)
		r.Departure.ActualTime.ChangeTimezone(location)

		r.Arrival.ScheduledTime.ChangeTimezone(location)
		r.Arrival.EstimatedTime.ChangeTimezone(location)
		r.Arrival.ActualTime.ChangeTimezone(location)

	}
}

func (ic *IATACodes) TimetableList(airport_iata string) ([]TimetableResponse, error) {

	d := TimetableRequest{}

	url := fmt.Sprintf("/v7/timetable?iata_code=%s&api_key=%s", airport_iata, ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []TimetableResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &d); err != nil {
		return []TimetableResponse{}, err
	}

	d.CorrectTimes()

	return d.Response, nil

}
