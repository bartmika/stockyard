package observation_average

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, u *ObservationAverage) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (*ObservationAverage, error)
	CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (bool, error)
	InsertOrGetByPrimaryKey(ctx context.Context, u *ObservationAverage) (*ObservationAverage, error)
	UpdateByPrimaryKey(ctx context.Context, u *ObservationAverage) error
	ListByFilter(ctx context.Context, filter *ObservationAverageFilter) ([]*ObservationAverage, error)
	CountByFilter(ctx context.Context, filter *ObservationAverageFilter) (uint64, error)
}
