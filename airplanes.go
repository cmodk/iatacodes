package iatacodes

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type MaybeInt int

func (mi *MaybeInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var err error
	var i int

	if b[0] == '"' {
		i, err = strconv.Atoi(strings.Replace(string(b), "\"", "", -1))
		if err != nil {
			return err
		}
	} else {
		log.Println("Custom unmarshal not string: %s\n", string(b))
	}

	*mi = MaybeInt(i)

	return nil
}

type AirlineTime time.Time

func (at *AirlineTime) UnmarshalJSON(b []byte) error {

	t, err := time.Parse(time.RFC3339, strings.Replace(string(b), "\"", "", -1))
	if err == nil {
		*at = AirlineTime(t)
	} else {
		log.Printf("Could not parse '%s'\n", string(b))
	}

	return nil
}

type Airplane struct {
	RegNumber    string      `json:"reg_number"`
	ICAOHex      string      `json:"icao_hex"`
	AircraftIATA string      `json:"aircraft_iata"`
	AircraftCode string      `json:"aircraft_code"`
	AircraftType string      `json:"aircraft_type"`
	EnginesType  string      `json:"engines_type"`
	EnginesCount MaybeInt    `json:"engines_count"`
	FirstFlight  AirlineTime `json:"first_flight"`
}

func (ic *IATACodes) AirplaneList() ([]Airplane, error) {

	d := struct {
		Request  ICRequest  `json:"request"`
		Response []Airplane `json:"response"`
	}{}

	url := fmt.Sprintf("/v6/airplanes?&api_key=%s&page=0&limit=100000", ic.key)

	resp, err := ic.sh.Get(url)
	if err != nil {
		return []Airplane{}, err
	}

	if ic.debug {
		log.Println(resp)
	}

	if err := json.Unmarshal([]byte(resp), &d); err != nil {
		return []Airplane{}, err
	}

	return d.Response, nil

}
