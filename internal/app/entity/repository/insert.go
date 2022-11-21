package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/entity"
)

func (r *EntityRepoImpl) Insert(ctx context.Context, m *domain.Entity) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sql, args, err := psql.
		Insert("entities").
		Columns(
			"uuid",
			"name",
			"data_type",
			"meta",
		).
		Values(
			m.UUID,
			m.Name,
			m.DataType,
			m.Meta,
		).
		ToSql()
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Str("uuid", m.UUID).
			Str("name", m.Name).
			Int8("data_type", m.DataType).
			Str("meta", m.Meta).
			Msg("failed creating query")
		return err
	}

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Str("uuid", m.UUID).
			Str("name", m.Name).
			Int8("data_type", m.DataType).
			Str("meta", m.Meta).
			Msg("prepare context error")
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	if err != nil {
		r.logger.Error().Err(err).Caller().
			Str("uuid", m.UUID).
			Str("name", m.Name).
			Int8("data_type", m.DataType).
			Str("meta", m.Meta).
			Msg("exec context error")
	}
	return err
}
