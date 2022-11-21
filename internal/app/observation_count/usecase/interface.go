package usecase

import (
	"context"

	odomain "github.com/bartmika/stockyard/internal/domain/observation_count"
	osum "github.com/bartmika/stockyard/internal/domain/observation_count"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

// Usecase Provides interface for the observation count use cases.
type Usecase interface {
	Insert(ctx context.Context, e *osum.ObservationCount) (ee *osum.ObservationCount, err error)
	ListAndCountByFilter(ctx context.Context, ef *osum.ObservationCountFilter) ([]*osum.ObservationCount, uint64, error)
}

type observationCountUsecase struct {
	Time                     timep.Provider
	UUID                     uuid.Provider
	ObservationCountRepo odomain.Repository
}

// NewObservationCountUsecase Constructor function for the `ObservationCountUsecase` implementation.
func NewObservationCountUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	o odomain.Repository,

) *observationCountUsecase {
	return &observationCountUsecase{
		Time:                     tp,
		UUID:                     uuidp,
		ObservationCountRepo: o,
	}
}
