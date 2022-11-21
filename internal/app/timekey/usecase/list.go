package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/timekey"
)

func (uc timekeyUsecase) ListAndCountByFilter(ctx context.Context, tkf *timekey.TimeKeyFilter) ([]*timekey.TimeKey, uint64, error) {
	arrCh := make(chan []*timekey.TimeKey)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.TimeKeyRepo.ListByFilter(ctx, tkf)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.TimeKeyRepo.CountByFilter(ctx, tkf)
		if err != nil {
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	return arr, count, nil
}
