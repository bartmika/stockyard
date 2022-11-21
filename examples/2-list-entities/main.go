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
// instance on your machine and list all the entities that you previously
// submitted.
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
	//// List all our entites.
	////

	// Generate the filter that you want to sort by.
	dto := &dtos.EntityFilterRequestDTO{
		SortOrder: "ASC",
		SortField: "id",
		Offset:    0,
		Limit:     1_000_000,
		IDs:       []uint64{},
	}

	// Execute the remote call.
	entities, err := client.ListEntities(dto)
	if err != nil {
		log.Fatal(err)
	}

	////
	//// View our results.
	////

	// Print the results.
	if entities.Count > 0 {
		for _, entity := range entities.Results {
			log.Println("id", entity.ID)
			log.Println("uuid", entity.UUID)
			log.Println("name", entity.Name)
			log.Println("data_type", entity.DataType)
			log.Println("meta", entity.Meta)
		}
	}
}
