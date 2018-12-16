package iatacodes

import (
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
	return &t
}
