package observation

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, u *Observation) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (*Observation, error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (bool, error)
	ListByFilter(ctx context.Context, filter *ObservationFilter) ([]*Observation, error)
	CountByFilter(ctx context.Context, filter *ObservationFilter) (uint64, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) error
}
