package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/observation"
)

func (uc observationUsecase) ListAndCountByFilter(ctx context.Context, ef *observation.ObservationFilter) ([]*observation.Observation, uint64, error) {
	arrCh := make(chan []*observation.Observation)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.ObservationRepo.ListByFilter(ctx, ef)
		if err != nil {
			uc.Logger.Warn().Err(err).Caller().Msg("database error")
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.ObservationRepo.CountByFilter(ctx, ef)
		if err != nil {
			uc.Logger.Warn().Err(err).Caller().Msg("database error")
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	if uc.HasAnalyzer {
		for _, o := range arr {
			uc.KMutex.Lockf("entity-id-%d", o.EntityID)
			defer uc.KMutex.Unlockf("entity-id-%d", o.EntityID)
		}
	}

	return arr, count, nil
}
