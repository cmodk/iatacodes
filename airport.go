package iatacodes

import (
	"encoding/json"
	"fmt"
)

type Airport struct {
	IATA         string  `json:"code"`
	Name         string  `json:"name"`
	ICAO         string  `json:"icao"`
	Longitude    float64 `json:"lng"`
	Latitude     float64 `json:"lat"`
	Country      string  `json:"country"`
	IsRailRoad   int     `json:"is_rail_road"`
	IsBusStation int     `json:"is_bus_station"`
}

func (ic *IATACodes) AirportList() ([]Airport, error) {

	d := struct {
		Request  ICRequest `json:"request"`
		Response []Airport `json:"response"`
	}{}

	url := fmt.Sprintf("/v6/airports?&api_key=%s", ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []Airport{}, err
	}

	if err := json.Unmarshal([]byte(resp), &d); err != nil {
		return []Airport{}, err
	}

	return d.Response, nil

}
