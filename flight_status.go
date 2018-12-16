package iatacodes

import (
	"encoding/json"
	"fmt"
)

type ICRequest struct {
	Language string `json:"lang"`
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

	url := fmt.Sprintf("/v7/routes?limit=10&api_key=%s", ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []RouteResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &fs); err != nil {
		return []RouteResponse{}, err
	}

	return fs.Response, nil

}

func (ic *IATACodes) TimetableList(airport_iata string) ([]TimetableResponse, error) {

	d := struct {
		Request  ICRequest           `json:"request"`
		Response []TimetableResponse `json:"response"`
	}{}

	url := fmt.Sprintf("/v7/timetable?limit=10&iata_code=%s&api_key=%s", airport_iata, ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []TimetableResponse{}, err
	}

	if err := json.Unmarshal([]byte(resp), &d); err != nil {
		return []TimetableResponse{}, err
	}

	return d.Response, nil

}
