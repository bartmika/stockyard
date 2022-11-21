package main

import (
	"fmt"
	"log"
	"time"

	rpc_client "github.com/bartmika/stockyard/pkg/rpc"
)

// DESCRIPTION:
// The purpose of this application is to connect to a running `stockyard` server
// instance on your machine and delete a previously submitted entity.
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
	deviceID := uint64(1)

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// Run the entry deletion.
	////

	// Execute the remote call.
	if err := client.DeleteEntityByPrimaryKey(deviceID); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully deleted an entity")
}
