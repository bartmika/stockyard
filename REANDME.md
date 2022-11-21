# stockyard
**IN PROGRESS - USE AT YOUR OWN RISK**
Simple distributed *time-series data* storage written in Go and powered by Postgres using the [Citus extension](https://github.com/citusdata/citus).

## Features:

* Runs in background accepting API calls.
* Make API calls from your Golang application over network or localhost.
* Horizontally scalable database if you know [Citus extension](https://github.com/citusdata/citus).
* Easy to setup and manage if you know [`docker`](https://www.docker.com).

## Install via Docker

Go to your workdir:

```shell
cd /your/project/directory
```

Get the latest Compose file:

```shell
curl https://raw.githubusercontent.com/xxx/docker-compose.yml > docker-compose.yml
```

Get the latest `env_file`:

```shell
curl https://raw.githubusercontent.com/xxx/.env > .env
```

Afterwords tweek the `.env` file, here are a few helpful comments:

* `STOCKYARD_DB_USER` - The admin user's username of the your database.
* `STOCKYARD_DB_PASSWORD` - The admin user's password of your database.
* `STOCKYARD_DB_NAME` - The database name. Recommended: `stockyard_db`.
* `STOCKYARD_APP_SECRET_KEY` - *(Optional)* The secret key to store and be used internally. *Ignore for now.*
* `STOCKYARD_HAS_ANALYZER` - The boolean value which controls whether to enable the analyzer to run on ever observation insert or delete.
* `PGADMIN_DEFAULT_EMAIL` - *(Optional)* The `pgadmin` administrators account email if you want to use `pgadmin`.
* `PGADMIN_DEFAULT_PASSWORD` - *(Optional)* The `pgadmin` administrators account password if you want to use `pgadmin`.

Tweak the `docker-compose.yml` file there according to your needs

```shell
$EDITOR ./docker-compose.yml
```

Start the server.

```shell
docker-compose -p stockyard -f docker-compose.yml up
```

## Usage

### 1. Create Entity
An **entity** represents a device, instrument or some grouping of time-series data. The following sample code demonstrates how to **create** two IoT instruments and add the accompanying readings into it.

There are two types of entities:

* **EntityObservationDataType** - Allows you to store data in the **Observation** model.
* **EntityTimeKeyDataType** - Allows you to store data in the **TimeKey** model.

```go
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
```


### 2. List Entities
Here is an example of how to list entities:

```go
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
```

### 3. Create Observation
An **observation** is essentially our time-series datum that belongs to an **entity**. Here is a coding example:

```go
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
```

### 4. List Observation
Now that you have data in the database, here is how you get a filtered list:

```go
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
```

### 5. Insert TimeKey
A **TimeKey** model is used if you want to store a `string` data-type value instead of the `float64` data-type that you find an `observation`. Why would you use this? A **TimeKey** model would be ideal to use if you store files in a remote S3 server and you want to store the s3 keys to access them in a timed manner. Here is an example:

```go
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
```

### 6. List TimeKey(s)
Here is how you can list your timekeys:

```go
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
```

### 7. Delete TimeKey(s)

Here is how you delete timekey(s):

```go
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
// instance on your machine and delete all the timekey(s) that you previously
// submitted.
//
// HOW TO RUN:
// cd ./examples/7-delete-timekey
// go run main.go
//
// PRECONDITION:
// You need to have inserted an entity before running this code.

func main() {
	////
	//// Connect to the running stockyard server running in the background.
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
	//// Delete the timekey(s).
	////

	// Generate the filter that you want to sort by.
	dto := &dtos.TimeKeyFilterRequestDTO{
		EntityIDs:                   []uint64{2},
		TimestampGreaterThenOrEqual: time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC),
		TimestampLessThenOrEqual:    time.Date(2023, 01, 01, 0, 0, 0, 0, time.UTC).AddDate(1, 0, 0),
	}

	// Execute the remote call.
	err = client.DeleteTimeKeysByFilter(dto)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("successfully deleted")
}
```

## Contributing

Found a bug? Want a feature to improve your developer experience when dealing with `stockyard`? Please create an [issue](https://github.com/bartmika/stockyard/issues).

## License
Made with ❤️ by [Bartlomiej Mika](https://bartlomiejmika.com).   
The project is licensed under the [ISC License](LICENSE).

Resource used:

* [citusdata/citus](https://github.com/citusdata/citus) is the distributed PostgreSQL extension that this server was built upon.
* [Clean Architecture by Panayiotis Kritiotis](https://github.com/pkritiotis/go-climb-clean-architecture-example) was the architecture pattern used on this codebase.
