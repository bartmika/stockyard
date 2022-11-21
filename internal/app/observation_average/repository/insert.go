package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_average"
)

func (r *ObservationAverageRepoImpl) Insert(ctx context.Context, m *domain.ObservationAverage) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.
		Insert("observation_averages").
		Columns(
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
		Values(
			m.EntityID,
			m.Start,
			m.Finish,
			m.Day,
			m.Week,
			m.Month,
			m.Year,
			m.Frequency,
			m.Result,
		).
		ToSql()
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			// Str("meta", m.Meta).
			// Time("timestamp", m.Timestamp).
			// Float64("value", m.Value).
			Msg("failed creating query")
		return err
	}

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			// Str("meta", m.Meta).
			// Time("timestamp", m.Timestamp).
			// Float64("value", m.Value).
			Msg("prepare context error")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			// Str("meta", m.Meta).
			// Time("timestamp", m.Timestamp).
			// Float64("value", m.Value).
			Msg("exec context error")
	}
	return err
}
