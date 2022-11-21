package observation_analyzer_request

import (
	"context"
)

type Repository interface {
	Insert(ctx context.Context, u *ObservationAnalyzerRequest) error
	GetByPrimaryKey(ctx context.Context, entityID uint64, uuid string) (*ObservationAnalyzerRequest, error)
	ListAll(ctx context.Context) ([]*ObservationAnalyzerRequest, error)
	DeleteByPrimaryKey(ctx context.Context, entityID uint64, uuid string) error
}
