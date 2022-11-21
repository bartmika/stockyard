package main

import (
	"fmt"
	"log"
	"time"

	"github.com/bartmika/stockyard/pkg/dtos"
	rpc_client "github.com/bartmika/stockyard/pkg/rpc"
)

// DESCRIPTION:
// The purpose of this application is to connect to a running `stockyard` server
// instance on your machine and list all the timekey(s) that you previously
// submitted.
//
// HOW TO RUN:
// cd ./examples/6-list-timekey
// go run main.go
//
// PRECONDITION:
// You need to have inserted an entity before running this code.

func main() {
	////
	//// Connect to our running stockyard server.
	////

	// Sample data to use in our example code.
	ipAddress := "127.0.0.1"
	port := "8000"

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// List our data within a time range.
	////

	// Generate the filter that you want to sort by.
	dto := &dtos.TimeKeyFilterRequestDTO{
		EntityIDs:                   []uint64{2},
		TimestampGreaterThenOrEqual: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		TimestampLessThenOrEqual:    time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC).AddDate(1, 0, 0),
	}

	// Execute the remote call.
	res, err := client.ListTimeKeys(dto)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// View our results.
	////

	// Print the results.
	if res.Count > 0 {
		for _, tk := range res.Results {
			log.Println("entity_id", tk.EntityID)
			log.Println("timestamp", tk.Timestamp)
			log.Println("data_type", tk.Value)
			log.Println("meta", tk.Meta)
		}
	} else {
		log.Println("no results returned")
	}
}
