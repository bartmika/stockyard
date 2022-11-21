package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation"
)

func (r *ObservationRepoImpl) getBy(ctx context.Context, k *sq.And) (*domain.Observation, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select(
			"entity_id",
			"meta",
			"timestamp",
			"value",
		).
		From("observations").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		r.logger.Error().Err(err).Caller().Msgf("prepare context error for k: %v", k)
		return nil, err
	}
	defer stmt.Close()

	m := new(domain.Observation)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&m.EntityID,
		&m.Meta,
		&m.Timestamp,
		&m.Value,
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

func (dr *ObservationRepoImpl) GetByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (*domain.Observation, error) {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"timestamp": timestamp},
	}
	return dr.getBy(ctx, k)
}
