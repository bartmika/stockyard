package dtos

import "time"

type ObservationCountResponseDTO struct {
	EntityID  uint64    `json:"entity_id"`
	Start     time.Time `json:"start"`
	Finish    time.Time `json:"finish"`
	Day       int       `json:"day"`
	Week      int       `json:"week"`
	Month     int       `json:"month"`
	Year      int       `json:"year"`
	Frequency int8      `json:"frequency"`
	Result    float64   `json:"result"`
}

type ObservationCountFilterRequestDTO struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	Frequency                   int8      `json:"frequency"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}

type ObservationCountListResponseDTO struct {
	Results []*ObservationCountResponseDTO `json:"results"`
	Count   uint64                         `json:"count"`
}
