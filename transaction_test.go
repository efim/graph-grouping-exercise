package main

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestTransactionMarshall(t *testing.T) {
	testTime := time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC)
	sampleTransaction := Transaction{
		Value: 1234,
		Timestamp: UnixTime{
			Time: testTime,
		},
	}

	expected := fmt.Sprintf(`{"value":%d,"timestamp":%d}`, sampleTransaction.Value, sampleTransaction.Timestamp.Time.Unix())

	bytes, err := json.Marshal(sampleTransaction)
	if err != nil {
		t.Fatal(err)
	}
	str := string(bytes)
	if str != expected {
		t.Fatal("marshalled doesn't equal to expected: ", str, expected)
	}

}

func TestTransactionUnmarshall(t *testing.T) {
	unixTimestamp := 1136214245
	value := 1234
	input := fmt.Sprintf(`{"value":%d,"timestamp":%d}`, value, unixTimestamp)

	expected := Transaction{
		Value: value,
		Timestamp: UnixTime{
			Time: time.Unix(int64(unixTimestamp), 0),
		},
	}

	var unmarshalled Transaction
	err := json.Unmarshal([]byte(input), &unmarshalled)
	if err != nil {
		t.Fatal(err)
	}

	if unmarshalled != expected {
		t.Fatal("unmarshalled doesn't equal expected: ", unmarshalled, expected)
	}
}
