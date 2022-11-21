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
// instance on your machine and create an entity.

func main() {
	////
	//// Connect to our running stockyard server.
	////

	// Sample data to use in our example code.
	ipAddress := "127.0.0.1"
	port := "8000"
	deviceName1 := "temperate-sensor-1" // Give your entity any unique name you like.
	deviceName2 := "backyard-birdfeeder-phototimer"

	// Connect to a running server from this appolication.
	applicationAddress := fmt.Sprintf("%s:%s", ipAddress, port)
	client, err := rpc_client.NewClient(applicationAddress, 3, 15*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// Create a `temperature sensor` collection.
	////

	// Execute the remote call.
	entity, err := client.InsertEntity(deviceName1, dtos.EntityObservationDataType, "")
	if err != nil {
		log.Fatal(err)
	}

	if entity == nil {
		log.Fatal("error as nothing was returned")
	}

	// See the results.
	log.Println("id", entity.ID)
	log.Println("uuid", entity.UUID)
	log.Println("name", entity.Name)
	log.Println("data_type", entity.DataType)
	log.Println("meta", entity.Meta)

	////
	//// Create a `phototimer` collection.
	////

	// Execute the remote call.
	entity, err = client.InsertEntity(deviceName2, dtos.EntityTimeKeyDataType, "")
	if err != nil {
		log.Fatal(err)
	}

	if entity == nil {
		log.Fatal("error as nothing was returned")
	}

	// See the results.
	log.Println("id", entity.ID)
	log.Println("uuid", entity.UUID)
	log.Println("name", entity.Name)
	log.Println("data_type", entity.DataType)
	log.Println("meta", entity.Meta)
}
