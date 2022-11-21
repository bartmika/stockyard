package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/entity"
)

func (uc entityUsecase) Insert(ctx context.Context, e *entity.Entity) (*entity.Entity, error) {
	e.UUID = uc.UUID.NewUUID()
	if err := uc.EntityRepo.Insert(ctx, e); err != nil {
		return nil, err
	}
	e, err := uc.EntityRepo.GetByUUID(ctx, e.UUID)
	if err != nil {
		return nil, err
	}
	return e, nil
}
