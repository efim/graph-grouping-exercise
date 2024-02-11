package main

import (
	"encoding/json"
	"io/ioutil"
	"reflect"
	"testing"
	"time"
)

func TestGrouping(t *testing.T) {
	filepath := "test-data/example-transactions-1.json"
	fileContent, err := ioutil.ReadFile(filepath)

	if err != nil {
		t.Fatal(err)
	}

	var inputTransactions []*Transaction
	err = json.Unmarshal(fileContent, &inputTransactions)

	if err != nil {
		t.Fatal(err)
	}

	expectedResultsFilepath := "test-data/example-result-1.json"
	expectedResultsFileContent, err := ioutil.ReadFile(expectedResultsFilepath)

	if err != nil {
		t.Fatal(err)
	}

	var expectedTransactions []*Transaction
	err = json.Unmarshal(expectedResultsFileContent, &expectedTransactions)

	if err != nil {
		t.Fatal(err)
	}

	grouppedByDay, err := GroupTransactions(inputTransactions, Day)
	if err != nil {
		t.Error(err)
	}

	if len(grouppedByDay) != len(expectedTransactions) {
		t.Error("unequal lengths of groupped and expected: ", grouppedByDay, expectedTransactions)
	}

	for i, v := range grouppedByDay {
		expected := expectedTransactions[i]
		if !reflect.DeepEqual(v, expected) {
			t.Errorf("slices differ at index %d : %s not equal to expected %s", i, v, expected)
		}
	}
}

func TestTruncateToHour(t *testing.T) {
	testTime := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	truncated, err := Truncate(testTime, Hour)

	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2006, time.January, 2, 15, 0, 0, 0, time.UTC)
	if truncated != expected {
		t.Fatal("truncated to hour is not equal to expected: ", truncated, expected)
	}

}

func TestTruncateToDay(t *testing.T) {
	testTime := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	truncated, err := Truncate(testTime, Day)

	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2006, time.January, 2, 0, 0, 0, 0, time.UTC)
	if truncated != expected {
		t.Fatal("truncated to day is not equal to expected: ", truncated, expected)
	}

}
func TestTruncateToWeek(t *testing.T) {
	testTime := time.Date(2024, time.February, 14, 15, 4, 5, 0, time.UTC)
	truncated, err := Truncate(testTime, Week)

	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2024, time.February, 12, 0, 0, 0, 0, time.UTC)
	if truncated != expected {
		t.Fatal("truncated to week is not equal to expected: ", truncated, expected)
	}

}
func TestTruncateToMonth(t *testing.T) {
	testTime := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	truncated, err := Truncate(testTime, Month)

	if err != nil {
		t.Fatal(err)
	}

	expected := time.Date(2006, time.January, 1, 0, 0, 0, 0, time.UTC)
	if truncated != expected {
		t.Fatal("truncated to month is not equal to expected: ", truncated, expected)
	}

}
