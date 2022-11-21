package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (r *ObservationAnalyzerRequestRepoImpl) getBy(ctx context.Context, k *sq.And) (*domain.ObservationAnalyzerRequest, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select(
			"entity_id",
			"uuid",
			"timestamp",
			"type",
			"observation",
		).
		From("observation_analyzer_requests").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		r.logger.Error().Err(err).Caller().Msgf("prepare context error for k: %v", k)
		return nil, err
	}
	defer stmt.Close()

	var obin []byte
	m := new(domain.ObservationAnalyzerRequest)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&m.EntityID,
		&m.UUID,
		&m.Timestamp,
		&m.Type,
		&obin,
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

	if err := json.Unmarshal(obin, &m.Observation); err != nil {
		return nil, err
	}

	return m, nil
}

func (dr *ObservationAnalyzerRequestRepoImpl) GetByPrimaryKey(ctx context.Context, entityID uint64, uuid string) (*domain.ObservationAnalyzerRequest, error) {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"uuid": uuid},
	}
	return dr.getBy(ctx, k)
}
