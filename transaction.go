package main

import (
	"encoding/json"
	"fmt"
	"time"
)

// working with Timestamp with expectation of UTC
type Transaction struct {
	Value     int      `json:"value"`
	Timestamp UnixTime `json:"timestamp"`
}

func (t *Transaction) String() string {
	return fmt.Sprint(*t)
}

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
