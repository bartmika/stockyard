package usecase

import (
	"context"

	osum "github.com/bartmika/stockyard/internal/domain/observation_average"
)

func (uc observationAverageUsecase) ListAndCountByFilter(ctx context.Context, ef *osum.ObservationAverageFilter) ([]*osum.ObservationAverage, uint64, error) {
	arrCh := make(chan []*osum.ObservationAverage)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.ObservationAverageRepo.ListByFilter(ctx, ef)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.ObservationAverageRepo.CountByFilter(ctx, ef)
		if err != nil {
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	return arr, count, nil
}
