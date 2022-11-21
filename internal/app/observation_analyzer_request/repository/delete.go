package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (r *ObservationAnalyzerRequestRepoImpl) deleteBy(ctx context.Context, k *sq.And) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)

	sql, args, err := psql.
		Delete("observation_analyzer_requests").
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

func (r *ObservationAnalyzerRequestRepoImpl) DeleteByPrimaryKey(ctx context.Context, entityID uint64, uuid string) error {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"uuid": uuid},
	}
	return r.deleteBy(ctx, k)
}
