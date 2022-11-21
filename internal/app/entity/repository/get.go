package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/entity"
)

func (r *EntityRepoImpl) getBy(ctx context.Context, k *sq.And) (*domain.Entity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select(
			"id",
			"uuid",
			"name",
			"data_type",
			"meta",
		).
		From("entities").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		r.logger.Error().Err(err).Caller().Msgf("prepare context error for k: %v", k)
		return nil, err
	}
	defer stmt.Close()

	m := new(domain.Entity)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&m.ID,
		&m.UUID,
		&m.Name,
		&m.DataType,
		&m.Meta,
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

func (dr *EntityRepoImpl) GetByID(ctx context.Context, id uint64) (*domain.Entity, error) {
	k := &sq.And{
		sq.Eq{"id": id},
	}
	return dr.getBy(ctx, k)
}

func (dr *EntityRepoImpl) GetByUUID(ctx context.Context, uid string) (*domain.Entity, error) {
	k := &sq.And{
		sq.Eq{"uuid": uid},
	}
	return dr.getBy(ctx, k)
}
