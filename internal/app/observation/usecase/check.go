package usecase

import (
	"context"
	"time"
)

func (uc observationUsecase) CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (bool, error) {
	doesExist, err := uc.ObservationRepo.CheckIfExistsByPrimaryKey(ctx, entityID, timestamp)
	if err != nil {
		uc.Logger.Error().Err(err).Caller().Msg("database error")
		return false, err
	}
	return doesExist, nil
}
