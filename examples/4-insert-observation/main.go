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
// instance on your machine and create an observation for a previously created
// entity.

func main() {
	////
	//// Connect to our running stockyard server.
	////

	// Sample data to use in our example code.
	ipAddress := "127.0.0.1"
	port := "8000"
	deviceID := uint64(1)
	deviceTimestamp := time.Date(2022, 1, 1, 1, 1, 1, 1, time.UTC)
	deviceValue := float64(300)

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// Create data.
	////

	// Execute the remote call.
	observationReq := &dtos.ObservationInsertRequestDTO{
		EntityID:  deviceID,
		Meta:      "",
		Timestamp: deviceTimestamp,
		Value:     deviceValue,
	}

	// Execute the remote call.
	observationRes, err := client.InsertObservation(observationReq)
	if err != nil {
		log.Fatal(err)
	}
	if observationRes == nil {
		log.Fatal("error as nothing was returned")
	}

	////
	//// View results.
	////

	// See the results.
	log.Println("entity_id", observationRes.EntityID)
	log.Println("timestamp", observationRes.Timestamp)
	log.Println("value", observationRes.Value)
	log.Println("meta", observationRes.Meta)
}
