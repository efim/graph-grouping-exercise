package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
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

	for _, tr := range inputTransactions {
		t.Log(*tr)
	}
}
