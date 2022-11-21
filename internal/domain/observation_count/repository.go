package observation_count

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, u *ObservationCount) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (*ObservationCount, error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (bool, error)
	InsertOrGetByPrimaryKey(ctx context.Context, u *ObservationCount) (*ObservationCount, error)
	UpdateByPrimaryKey(ctx context.Context, u *ObservationCount) error
	ListByFilter(ctx context.Context, filter *ObservationCountFilter) ([]*ObservationCount, error)
	CountByFilter(ctx context.Context, filter *ObservationCountFilter) (uint64, error)
}
