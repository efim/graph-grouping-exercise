package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func run(filename string, intervalName string) error {
	interval, err := ParseGroupingInterval(intervalName)
	if err != nil {
		return err
	}

	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	var inputTransactions []*Transaction
	err = json.Unmarshal(fileContent, &inputTransactions)

	if err != nil {
		return err
	}

	grouppedTransactions, err := GroupTransactions(inputTransactions, interval)
	if err != nil {
		return err
	}

	fmt.Println("Groupped transactions:")
	for _, transaction := range grouppedTransactions {
		fmt.Println(transaction)
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Specify filename to process and groupping interval")
		os.Exit(1)
	}

	filename := os.Args[1]
	fmt.Println("Will group transactions from file: ", filename)

	err := run(filename, os.Args[2])
	if err != nil {
		fmt.Println("Error groupping data: ", err)
		os.Exit(1)
	}
}
