package usecase

import (
	"context"

	"github.com/bartmika/stockyard/internal/domain/entity"
	edomain "github.com/bartmika/stockyard/internal/domain/entity"
	timep "github.com/bartmika/stockyard/internal/pkg/time"
	"github.com/bartmika/stockyard/internal/pkg/uuid"
)

// EntityUsecase Provides interface for the entity use cases.
type Usecase interface {
	Insert(ctx context.Context, e *entity.Entity) (ee *entity.Entity, err error)
	ListAndCountByFilter(ctx context.Context, ef *entity.EntityFilter) ([]*entity.Entity, uint64, error)
	Delete(ctx context.Context, entityID uint64) (err error)
}

type entityUsecase struct {
	Time       timep.Provider
	UUID       uuid.Provider
	EntityRepo edomain.Repository
}

// NewEntityUsecase Constructor function for the `UserUsecase` implementation.
func NewEntityUsecase(
	uuidp uuid.Provider,
	tp timep.Provider,
	e edomain.Repository,

) *entityUsecase {
	return &entityUsecase{
		Time:       tp,
		UUID:       uuidp,
		EntityRepo: e,
	}
}
