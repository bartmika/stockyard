package usecase

import (
	"context"
)

func (uc observationAnalyzerRequestUsecase) DeleteByPrimaryKey(ctx context.Context, entityID uint64, uuid string) error {
	if err := uc.ObservationAnalyzerRequestRepo.DeleteByPrimaryKey(ctx, entityID, uuid); err != nil {
		return err
	}
	return nil
}
