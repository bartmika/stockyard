package usecase

import (
	"context"

	osum "github.com/bartmika/stockyard/internal/domain/observation_average"
)

func (uc observationAverageUsecase) Insert(ctx context.Context, os *osum.ObservationAverage) (*osum.ObservationAverage, error) {
	if err := uc.ObservationAverageRepo.Insert(ctx, os); err != nil {
		return nil, err
	}
	os, err := uc.ObservationAverageRepo.GetByPrimaryKey(ctx, os.EntityID, os.Frequency, os.Start, os.Finish)
	if err != nil {
		return nil, err
	}
	return os, nil
}
