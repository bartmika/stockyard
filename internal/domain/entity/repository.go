package entity

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, u *Entity) error
	GetByID(ctx context.Context, id uint64) (*Entity, error)
	GetByUUID(ctx context.Context, uuid string) (*Entity, error)
	ListByFilter(ctx context.Context, filter *EntityFilter) ([]*Entity, error)
	CountByFilter(ctx context.Context, filter *EntityFilter) (uint64, error)
	// UpdateByID(ctx context.Context, u *Entity) error
	// CheckIfExistsByID(ctx context.Context, id string) (bool, error)
	// InsertOrUpdateByID(ctx context.Context, u *Entity) error
	DeleteByID(ctx context.Context, id uint64) error
}
