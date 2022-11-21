package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/timekey"
)

func (uc timekeyUsecase) DeleteByFilter(ctx context.Context, tkf *timekey.TimeKeyFilter) error {
	return uc.TimeKeyRepo.DeleteByFilter(ctx, tkf)
}
