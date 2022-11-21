package entity

const (
	EntityObservationDataType = 1
	EntityTimeKeyDataType     = 2
)

type Entity struct {
	ID       uint64 `json:"id"`        // ex: 1
	UUID     string `json:"uuid"`      // ex: xxxx-xxxx-xxxx-xxxx
	Name     string `json:"name"`      // ex: Temperature Sensor
	DataType int8   `json:"data_type"` // ex: 1=float64
	Meta     string `json:"meta"`      // ex: {"user_id":1,"device_id",1234}
}

type EntityFilter struct {
	SortOrder string   `json:"sort_order"`
	SortField string   `json:"sort_field"`
	Offset    uint64   `json:"offset"`
	Limit     uint64   `json:"limit"`
	IDs       []uint64 `json:"ids"`
	DataType  int8     `json:"data_type"`
}
