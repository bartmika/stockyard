package repository

import (
	"context"
	"database/sql"
	"time"

	sq "github.com/Masterminds/squirrel"
)

func (r *ObservationRepoImpl) checkIfExistsBy(ctx context.Context, k *sq.And) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(r.dbCache).PlaceholderFormat(sq.Dollar)
	sqlQuery, args, err := psql.
		Select("1").
		From("observations").
		Where(k).
		ToSql()

	stmt, err := r.db.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	exists := new(bool)
	err = stmt.QueryRowContext(ctx, args...).Scan(
		&exists,
	)
	if err != nil {
		// CASE 1 OF 2: Cannot find record with that email.
		if err == sql.ErrNoRows {
			return false, nil
		}
		// CASE 2 OF 2: All other errors.
		return false, err
	}

	return *exists, nil
}

func (dr *ObservationRepoImpl) CheckIfExistsByPrimaryKey(ctx context.Context, entityID uint64, timestamp time.Time) (bool, error) {
	k := &sq.And{
		sq.Eq{"entity_id": entityID},
		sq.Eq{"timestamp": timestamp},
	}
	return dr.checkIfExistsBy(ctx, k)
}
