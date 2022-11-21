package observation_analyzer_request

import (
	"time"

	d "github.com/bartmika/stockyard/internal/domain/observation"
)

const (
	ObservationAnalyzerRequestInsertType = 1
	ObservationAnalyzerRequestDeleteType = 2
)

type ObservationAnalyzerRequest struct {
	EntityID    uint64         `json:"entity_id"`   // ex: 1234
	UUID        string         `json:"uuid"`        // ex: xxxx-xxx-xxxx-xxxx
	Timestamp   time.Time      `json:"timestamp"`   // ex: 2022-10-09 15:57:52 -0400 -0400
	Type        int8           `json:"type"`        // ex: 1
	Observation *d.Observation `json:"observation"` // ex: { ... }
}

type ObservationAnalyzerRequestFilter struct {
	EntityIDs                   []uint64  `json:"entity_ids"`
	TimestampGreaterThen        time.Time `json:"timestamp_gt,omitempty"`
	TimestampGreaterThenOrEqual time.Time `json:"timestamp_gte,omitempty"`
	TimestampLessThen           time.Time `json:"timestamp_lt,omitempty"`
	TimestampLessThenOrEqual    time.Time `json:"timestamp_lte,omitempty"`
}
