package usecase

import (
	"context"

	odomain "github.com/bartmika/stockyard/internal/domain/observation_summation"
	osum "github.com/bartmika/stockyard/internal/domain/observation_summation"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

// Usecase Provides interface for the observation summation use cases.
type Usecase interface {
	Insert(ctx context.Context, e *osum.ObservationSummation) (ee *osum.ObservationSummation, err error)
	ListAndCountByFilter(ctx context.Context, ef *osum.ObservationSummationFilter) ([]*osum.ObservationSummation, uint64, error)
}

type observationSummationUsecase struct {
	Time                     timep.Provider
	UUID                     uuid.Provider
	ObservationSummationRepo odomain.Repository
}

// NewObservationSummationUsecase Constructor function for the `ObservationSummationUsecase` implementation.
func NewObservationSummationUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o odomain.Repository,

) *observationSummationUsecase {
	return &observationSummationUsecase{
		Time:                     tp,
		UUID:                     uuidp,
		ObservationSummationRepo: o,
	}
}
