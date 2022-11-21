package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (r *ObservationRepoImpl) deleteBy(ctx context.Context, k *sq.And) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Delete("observations").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	return err
}

func (r *ObservationRepoImpl) DeleteByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) error {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"timestamp": timestamp},
	}
	return r.deleteBy(ctx, k)
}
