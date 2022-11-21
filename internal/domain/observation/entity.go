package observation

import "time"

type Observation struct {
	EntityID  uint64    `json:"entity_id"` // ex: xxxx-xxx-xxxx-xxxx
	Meta      string    `json:"meta"`      // ex: {"type":"heart_rate"}
	Timestamp time.Time `json:"timestamp"` // ex: 2022-10-09 15:57:52 -0400 -0400
	Value     float64   `json:"value"`     // ex: 123.456
}

type ObservationFilter struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}
