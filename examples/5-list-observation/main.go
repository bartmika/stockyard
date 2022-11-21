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
// instance on your machine and list all the observations that you previously
// submitted.
//
// PRECONDITION:
// You need to have inserted an entity and observation(s) before running this code.

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
	//// List our time-series data within a time range.
	////

	// Generate the filter that you want to sort by.
	req := &dtos.ObservationFilterRequestDTO{
		EntityIDs:                   []uint64{1},
		TimestampGreaterThenOrEqual: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		TimestampLessThenOrEqual:    time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC).AddDate(1, 0, 0),
	}

	// Execute the remote call.
	res, err := client.ListObservations(req)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// View our results.
	////

	// Print the results.
	if res.Count > 0 {
		for _, observation := range res.Results {
			log.Println("entity_id", observation.EntityID)
			log.Println("meta", observation.Meta)
			log.Println("timestamp", observation.Timestamp)
			log.Println("value", observation.Value)
		}
	}
}
