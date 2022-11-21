package usecase

import (
	"context"

	oardomain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (uc observationAnalyzerRequestUsecase) Insert(ctx context.Context, req *oardomain.ObservationAnalyzerRequest) (*oardomain.ObservationAnalyzerRequest, error) {
	if err := uc.ObservationAnalyzerRequestRepo.Insert(ctx, req); err != nil {
		return nil, err
	}
	o, err := uc.ObservationAnalyzerRequestRepo.GetByPrimaryKey(ctx, req.EntityID, req.UUID)
	if err != nil {
		return nil, err
	}
	return o, nil
}
