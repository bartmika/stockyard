package main

import (
	"fmt"
	"log"
	"time"

	model "github.com/bartmika/stockyard/internal/domain/observation_summation"
	"github.com/bartmika/stockyard/pkg/dtos"
	rpc_client "github.com/bartmika/stockyard/pkg/rpc"
)

// DESCRIPTION:
// The purpose of this application is to connect to a running `stockyard` server
// instance on your machine and list all the timekey(s) that you previously
// submitted.
//
// HOW TO RUN:
// cd ./examples/10-list-observation-summation
// go run main.go
//
// PRECONDITION:
// You need to have inserted an entity with multiple observation before running this code.
//    cd ../1-insert-entity; go run main.go;
//    cd ../4-insert-observation; go run main.go;

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
	dto := &dtos.ObservationSummationFilterRequestDTO{
		EntityIDs:               []uint64{1},
		Frequency:               model.ObservationSummationDayFrequency,
		StartGreaterThenOrEqual: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		FinishLessThenOrEqual:   time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC).AddDate(1, 0, 0),
	}

	// Execute the remote call.
	res, err := client.ListObservationSummations(dto)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// View our results.
	////

	// Print the results.
	if res.Count > 0 {
		log.Println("the `summation` results from the server are as follows:")
		for _, tk := range res.Results {
			fmt.Println("entity_id", tk.EntityID)
			fmt.Println("start", tk.Start)
			fmt.Println("finish", tk.Finish)
			// fmt.Println("day", tk.Day)     // Not important but useful if you need it.
			// fmt.Println("week", tk.Week)   // Not important but useful if you need it.
			// fmt.Println("month", tk.Month) // Not important but useful if you need it.
			// fmt.Println("year", tk.Year)   // Not important but useful if you need it.
			fmt.Println("frequency", tk.Frequency)
			fmt.Println("result", tk.Result)
			fmt.Println("")
		}
	} else {
		log.Println("no results returned")
	}
}
