package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_count"
)

func (r *ObservationCountRepoImpl) getBy(ctx context.Context, k *sq.And) (*domain.ObservationCount, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select(
			"entity_id",
			"start",
			"finish",
			"day",
			"week",
			"month",
			"year",
			"frequency",
			"result",
		).
		From("observation_counts").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		r.logger.Error().Err(err).Caller().Msgf("prepare context error for k: %v", k)
		return nil, err
	}
	defer stmt.Close()

	m := new(domain.ObservationCount)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&m.EntityID,
		&m.Start,
		&m.Finish,
		&m.Day,
		&m.Week,
		&m.Month,
		&m.Year,
		&m.Frequency,
		&m.Result,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return nil, nil
		}
		// CASE 2 OF 2: All other errors.
		r.logger.Error().Err(err).Caller().Msgf("query row context error for k: %v", k)
		return nil, err
	}

	return m, nil
}

func (dr *ObservationCountRepoImpl) GetByPrimaryKey(ctx context.Context, entityID uint64, frequency int8, start time.Time, finish time.Time) (*domain.ObservationCount, error) {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"frequency": frequency},
		sq.Eq{"start": start},
		sq.Eq{"finish": finish},
	}
	return dr.getBy(ctx, k)
}
