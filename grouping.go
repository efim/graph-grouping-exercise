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

func GroupTransactions(transactions []*Transaction, interval GroupingInterval) ([]*Transaction, error) {
	groupped := make([]*Transaction, 0)
	var truncatedTime *time.Time
	var latestTransaction *Transaction
	for index, transaction := range transactions {
		currentTransactionTruncatedTime, err := Truncate(transaction.Timestamp.Time, interval)
		if err != nil {
			return nil, err
		}
		if truncatedTime == nil {
			// first iteraction
			truncatedTime = &currentTransactionTruncatedTime
			latestTransaction = transaction
		}
		if currentTransactionTruncatedTime == *truncatedTime {
			// still processing same bucket
			// compare and keep latest transaction
			if latestTransaction.Timestamp.Time.Before(transaction.Timestamp.Time) {
				latestTransaction = transaction
			}
		} else {
			// closing old bucket
			truncatedTransaction := *latestTransaction
			truncatedTransaction.Timestamp.Time = *truncatedTime
			groupped = append(groupped, &truncatedTransaction)
			// starting new bucket
			truncatedTime = &currentTransactionTruncatedTime
			latestTransaction = transaction
		}

		if index == len(transactions)-1 {
			// processing the last transaction
			// 'closing old bucket' will not get evaluated
			// closing last bucket
			truncatedTransaction := *latestTransaction
			truncatedTransaction.Timestamp.Time = *truncatedTime
			groupped = append(groupped, &truncatedTransaction)
		}
	}
	return groupped, nil
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

func ParseGroupingInterval(input string) (GroupingInterval, error) {
	switch input {
	case string(Hour):
		return Hour, nil
	case string(Day):
		return Day, nil
	case string(Week):
		return Week, nil
	case string(Month):
		return Month, nil
	default:
		return "", fmt.Errorf("invalid grouping interval %s", input)
	}
}
