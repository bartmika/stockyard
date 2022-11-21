package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/timekey"
)

func (uc timekeyUsecase) Insert(ctx context.Context, tk *timekey.TimeKey) (*timekey.TimeKey, error) {
	if err := uc.TimeKeyRepo.Insert(ctx, tk); err != nil {
		return nil, err
	}
	tk, err := uc.TimeKeyRepo.GetByPrimaryKey(ctx, tk.EntityID, tk.Timestamp)
	if err != nil {
		return nil, err
	}
	return tk, nil
}
