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
// instance on your machine and create a timekey for a previously created
// entity.
//
// HOW TO RUN:
// cd ./examples/5-insert-timekey
// go run main.go

func main() {
	////
	//// Connect to our running stockyard server.
	////

	// Sample data to use in our example code.
	ipAddress := "127.0.0.1"
	port := "8000"
	deviceID := uint64(2) // Note: Assuming `2` was your phototimer entity.
	value := "s3key/some/remote/location/dir/photo1.png"

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// Insert the data.
	////

	// Execute the remote call.
	req := &dtos.TimeKeyInsertRequestDTO{
		EntityID:  deviceID,
		Meta:      "",
		Timestamp: time.Now(),
		Value:     value,
	}

	// Execute the remote call.
	res, err := client.InsertTimeKey(req)
	if err != nil {
		log.Fatal(err)
	}
	if res == nil {
		log.Fatal("error as nothing was returned")
	}

	////
	//// View the result.
	////

	// See the results.
	log.Println("entity_id", res.EntityID)
	log.Println("timestamp", res.Timestamp)
	log.Println("value", res.Value)
	log.Println("meta", res.Meta)
}
