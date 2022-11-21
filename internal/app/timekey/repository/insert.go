package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/timekey"
)

func (r *TimeKeyRepoImpl) Insert(ctx context.Context, m *domain.TimeKey) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.
		Insert("timekeys").
		Columns(
			"entity_id",
			"meta",
			"timestamp",
			"value",
		).
		Values(
			m.EntityID,
			m.Meta,
			m.Timestamp,
			m.Value,
		).
		ToSql()
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("meta", m.Meta).
			Time("timestamp", m.Timestamp).
			Str("value", m.Value).
			Msg("failed creating query")
		return err
	}

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("meta", m.Meta).
			Time("timestamp", m.Timestamp).
			Str("value", m.Value).
			Msg("prepare context error")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Uint64("entity_id", m.EntityID).
			Str("meta", m.Meta).
			Time("timestamp", m.Timestamp).
			Str("value", m.Value).
			Msg("exec context error")
	}
	return err
}
