package usecase

import (
	"context"

	osum "github.com/bartmika/stockyard/internal/domain/observation_summation"
)

func (uc observationSummationUsecase) Insert(ctx context.Context, os *osum.ObservationSummation) (*osum.ObservationSummation, error) {
	if err := uc.ObservationSummationRepo.Insert(ctx, os); err != nil {
		return nil, err
	}
	os, err := uc.ObservationSummationRepo.GetByPrimaryKey(ctx, os.EntityID, os.Frequency, os.Start, os.Finish)
	if err != nil {
		return nil, err
	}
	return os, nil
}
