package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Transaction struct {
	Value     int      `json:"value"`
	Timestamp UnixTime `json:"timestamp"`
}

func (t *Transaction) String() string {
	return fmt.Sprint(*t)
}

// this should allow working with input files
// where timestamp is encoded as unix timestamp
type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}

	u.Time = time.Unix(timestamp, 0)
	return nil
}
func (u UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", (u.Time.Unix()))), nil
}
