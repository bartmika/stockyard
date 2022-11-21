package repository

import (
	"context"

	domain "github.com/bartmika/stockyard/internal/domain/observation_count"
)

func (r *ObservationCountRepoImpl) InsertOrGetByPrimaryKey(ctx context.Context, oc *domain.ObservationCount) (*domain.ObservationCount, error) {
	doesExist, err := r.CheckIfExistsByPrimaryKey(
		ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish,
	)
	if err != nil {
		return nil, err
	}

	if doesExist == true {
		return r.GetByPrimaryKey(ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish)
	}
	if err := r.Insert(ctx, oc); err != nil {
		return nil, err
	}
	return r.GetByPrimaryKey(ctx, oc.EntityID, oc.Frequency, oc.Start, oc.Finish)
}
