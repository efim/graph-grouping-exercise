package main

import (
	"fmt"
	"time"
)

type GroupingInterval string

const (
	Hour  GroupingInterval = "hour"
	Day                    = "day"
	Week                   = "week"
	Month                  = "month"
)

func GroupTransactions(transactions []*Transaction, interval GroupingInterval) (groupped []*Transaction) {
	fmt.Println("gropuing")
	return
}

func Truncate(t time.Time, interval GroupingInterval) (time.Time, error) {
	switch interval {
	case Hour:
		truncated := t.Truncate(time.Hour)
		return truncated, nil
	case Day:
		truncated := t.Truncate(time.Hour * 24)
		return truncated, nil
	case Week:
		dayOffset := int(t.Weekday()) - int(time.Monday)
		truncated := time.Date(t.Year(), t.Month(), t.Day()-dayOffset, 0, 0, 0, 0, t.Location())
		return truncated, nil
	case Month:
		truncated := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
		return truncated, nil
	default:
		return t, fmt.Errorf("error calculating ceiling for interval %s from timestamp %s", interval, t)
	}
}
