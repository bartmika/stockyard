package dtos

import "time"

type ObservationResponseDTO struct {
	EntityID  uint64    `json:"entity_id"` // ex: 123
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     float64   `json:"value"`     // ex: 123.456
}

type ObservationFilterRequestDTO struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}

type ObservationListResponseDTO struct {
	Results []*ObservationResponseDTO `json:"results"`
	Count   uint64                    `json:"count"`
}

type ObservationPrimaryKeyRequestDTO struct {
	EntityID  uint64    `json:"entity_id"` // ex: 123
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
}

type ObservationInsertRequestDTO struct {
	UUID      string    `json:"uuid"`
	EntityID  uint64    `json:"entity_id"` // ex: xxxx-xxx-xxxx-xxxx
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     float64   `json:"value"`     // ex: 123.456
}
