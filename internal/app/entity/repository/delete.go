package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (r *EntityRepoImpl) deleteBy(ctx context.Context, k *sq.And) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Delete("entities").
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

func (r *EntityRepoImpl) DeleteByID(ctx context.Context, id uint64) error {
	k := &sq.And{
		sq.Eq{"id": id},
	}
	return r.deleteBy(ctx, k)
}
