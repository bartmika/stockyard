package usecase

import (
	"context"
)

func (uc entityUsecase) Delete(ctx context.Context, entityID uint64) error {
	if err := uc.EntityRepo.DeleteByID(ctx, entityID); err != nil {
		return err
	}
	return nil
}
