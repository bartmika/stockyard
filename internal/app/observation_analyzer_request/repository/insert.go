package repository

import (
	"context"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

func (r *ObservationAnalyzerRequestRepoImpl) Insert(ctx context.Context, m *domain.ObservationAnalyzerRequest) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	obin, err := json.Marshal(m.Observation)
	if err != nil {
		return err
	}

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.
		Insert("observation_analyzer_requests").
		Columns(
			"entity_id",
			"uuid",
			"timestamp",
			"type",
			"observation",
		).
		Values(
			m.EntityID,
			m.UUID,
			m.Timestamp,
			m.Type,
			obin,
		).
		ToSql()
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("uuid", m.UUID).
			Time("timestamp", m.Timestamp).
			Int8("type", m.Type).
			Str("observation", string(obin)).
			Msg("failed creating query")
		return err
	}

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("uuid", m.UUID).
			Time("timestamp", m.Timestamp).
			Int8("type", m.Type).
			Str("observation", string(obin)).
			Msg("prepare context error")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("uuid", m.UUID).
			Time("timestamp", m.Timestamp).
			Int8("type", m.Type).
			Str("observation", string(obin)).
			Msg("exec context error")
	}
	return err
}
