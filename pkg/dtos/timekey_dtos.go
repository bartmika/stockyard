package dtos

import "time"

type TimeKeyResponseDTO struct {
	EntityID  uint64    `json:"entity_id"` // ex: xxxx-xxx-xxxx-xxxx
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     string    `json:"value"`     // ex: "/filepath/locatino/1/2/3/example.txt"
}

type TimeKeyFilterRequestDTO struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}

type TimeKeyListResponseDTO struct {
	Results []*TimeKeyResponseDTO `json:"results"`
	Count   uint64                `json:"count"`
}

type TimeKeyInsertRequestDTO struct {
	UUID      string    `json:"uuid"`
	EntityID  uint64    `json:"entity_id"` // ex: xxxx-xxx-xxxx-xxxx
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     string    `json:"value"`     // ex: 123.456
}
