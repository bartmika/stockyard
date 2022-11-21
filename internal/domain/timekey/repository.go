package timekey

import (
	"context"
	"time"
)

type Repository interface {
	Insert(ctx context.Context, u *TimeKey) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (*TimeKey, error)
	ListByFilter(ctx context.Context, filter *TimeKeyFilter) ([]*TimeKey, error)
	CountByFilter(ctx context.Context, filter *TimeKeyFilter) (uint64, error)
	DeleteByFilter(ctx context.Context, filter *TimeKeyFilter) error
}
