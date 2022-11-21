package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/entity"
)

func (uc entityUsecase) ListAndCountByFilter(ctx context.Context, ef *entity.EntityFilter) ([]*entity.Entity, uint64, error) {
	arrCh := make(chan []*entity.Entity)
	countCh := make(chan uint64)

	go func() {
		arr, err := uc.EntityRepo.ListByFilter(ctx, ef)
		if err != nil {
			arrCh <- nil
			return
		}
		arrCh <- arr[:]
	}()

	go func() {
		count, err := uc.EntityRepo.CountByFilter(ctx, ef)
		if err != nil {
			countCh <- 0
			return
		}
		countCh <- count
	}()

	arr, count := <-arrCh, <-countCh

	return arr, count, nil
}
