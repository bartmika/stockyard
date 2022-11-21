package usecase

import (
	"context"

	oavg_d "github.com/bartmika/stockyard/internal/domain/observation_average"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

// Usecase Provides interface for the observation average use cases.
type Usecase interface {
	Insert(ctx context.Context, e *oavg_d.ObservationAverage) (ee *oavg_d.ObservationAverage, err error)
	ListAndCountByFilter(ctx context.Context, ef *oavg_d.ObservationAverageFilter) ([]*oavg_d.ObservationAverage, uint64, error)
}

type observationAverageUsecase struct {
	Time                   timep.Provider
	UUID                   uuid.Provider
	ObservationAverageRepo oavg_d.Repository
}

// NewObservationAverageUsecase Constructor function for the `ObservationAverageUsecase` implementation.
func NewObservationAverageUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o oavg_d.Repository,

) *observationAverageUsecase {
	return &observationAverageUsecase{
		Time:                   tp,
		UUID:                   uuidp,
		ObservationAverageRepo: o,
	}
}
