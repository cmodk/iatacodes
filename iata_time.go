package iatacodes

import (
	"log"
	"time"
)

type IATATime time.Time

func (it *IATATime) UnmarshalJSON(b []byte) error {

	rfc3339 := string(b[1:len(b)-1]) + "Z"

	t, err := time.Parse(time.RFC3339, rfc3339)
	if err != nil {
		return err
	}

	*it = IATATime(t)

	return nil
}

func (it *IATATime) GetTime() time.Time {
	return time.Time(*it)
}

func (it *IATATime) GetTimePtr() *time.Time {
	t := time.Time(*it)
	if t.IsZero() {
		return nil
	}
	return &t
}

func (it *IATATime) ChangeTimezone(loc *time.Location) {

	if it.GetTime().IsZero() {
		if debug {
			log.Printf("Zero time, ignore..\n")
		}
		return
	}

	t := it.GetTime()
	if debug {
		log.Printf("Original time %s to %s zone\n", t.Format(time.RFC3339), loc.String())
	}
	new_time := time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second(),
		t.Nanosecond(),
		loc).UTC()

	if debug {
		log.Printf("Correct time %s in %s zone\n", new_time.Format(time.RFC3339), loc.String())
	}

	new_iata_time := IATATime(new_time)

	*it = new_iata_time
}
