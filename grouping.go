package main

import (
	"fmt"
)

type GroupingInterval string

const (
	Hour GroupingInterval = "hour"
	Day = "day"
	Week = "week"
	Month = "month"
)

func GroupTransactions(transactions []*Transaction, interval GroupingInterval) (groupped []*Transaction) {
	fmt.Println("gropuing")
	return
}
