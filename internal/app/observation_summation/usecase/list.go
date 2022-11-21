package usecase

import (
	"context"

	osum "github.com/bartmika/stockyard/internal/domain/observation_summation"
)

func (uc observationSummationUsecase) ListAndCountByFilter(ctx context.Context, ef *osum.ObservationSummationFilter) ([]*osum.ObservationSummation, uint64, error) {
	arrCh := make(chan []*osum.ObservationSummation)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.ObservationSummationRepo.ListByFilter(ctx, ef)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.ObservationSummationRepo.CountByFilter(ctx, ef)
		if err != nil {
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	return arr, count, nil
}
