package observation_summation

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, u *ObservationSummation) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (*ObservationSummation, error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (bool, error)
	InsertOrGetByPrimaryKey(ctx context.Context, u *ObservationSummation) (*ObservationSummation, error)
	UpdateByPrimaryKey(ctx context.Context, u *ObservationSummation) error
	ListByFilter(ctx context.Context, filter *ObservationSummationFilter) ([]*ObservationSummation, error)
	CountByFilter(ctx context.Context, filter *ObservationSummationFilter) (uint64, error)
}
