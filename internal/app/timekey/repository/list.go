package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/timekey"
)

func (s *TimeKeyRepoImpl) getWhereKeysByFilter(f *domain.TimeKeyFilter) sq.And {
	// Apply specific 'where' keys to apply.
	k := sq.And{}

	if len(f.EntityIDs) > 0 {
		entityIDsKey := sq.Or{}
		for _, entityID := range f.EntityIDs {
			entityIDsKey = append(entityIDsKey, sq.Eq{"entity_id": entityID})
		}
		k = append(k, entityIDsKey)
	}

	if !f.TimestampGreaterThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Gt{"timestamp": f.TimestampGreaterThenOrEqual},
			sq.Eq{"timestamp": f.TimestampGreaterThenOrEqual},
		})
	}
	if !f.TimestampGreaterThen.IsZero() {
		k = append(k, sq.Gt{"timestamp": f.TimestampGreaterThen})
	}
	if !f.TimestampLessThen.IsZero() {
		k = append(k, sq.Lt{"timestamp": f.TimestampLessThen})
	}
	if !f.TimestampLessThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Lt{"timestamp": f.TimestampLessThenOrEqual},
			sq.Eq{"timestamp": f.TimestampLessThenOrEqual},
		})
	}

	return k
}

func (s *TimeKeyRepoImpl) ListByFilter(ctx context.Context, f *domain.TimeKeyFilter) ([]*domain.TimeKey, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
		"entity_id",
		"meta",
		"timestamp",
		"value",
	).From("timekeys")
	k := s.getWhereKeysByFilter(f)

	rds = rds.Where(k)

	rds = rds.OrderBy("timestamp" + " " + "ASC")

	// Note:
	// (1) https://ivopereira.net/efficient-pagination-dont-use-offset-limit
	// (2) https://github.com/Masterminds/squirrel/blob/def598cbb358368fbfc3f6a9a914699a36846992/select_test.go#L41

	// rds = rds.Offset(f.Offset).Suffix("FETCH FIRST ? ROWS ONLY", f.Limit)

	// Build the SQL statement and the accomponing arguments.
	sql, args, err := rds.ToSql()

	// // For debugging purposes only.
	// log.Println("sql:", sql)
	// log.Println("args:", args)
	// log.Println("err:", err)

	stmt, err := s.db.Prepare(sql)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("failed prepare with k: %v", k)
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("failed query context with k: %v", k)
		return nil, err
	}

	var arr []*domain.TimeKey
	defer rows.Close()
	for rows.Next() {
		m := new(domain.TimeKey)
		err := rows.Scan(
			&m.EntityID,
			&m.Meta,
			&m.Timestamp,
			&m.Value,
		)
		if err != nil {
			s.logger.Error().Err(err).Caller().Msgf("database scan error with k: %v", k)
			return nil, err
		}
		arr = append(arr, m)
	}
	err = rows.Err()
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("database with k: %v", k)
		return nil, err
	}

	if arr == nil {
		return []*domain.TimeKey{}, nil
	}
	return arr, err
}

func (s *TimeKeyRepoImpl) CountByFilter(ctx context.Context, f *domain.TimeKeyFilter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// The result we are looking for.
	var count uint64

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	submissionCount := psql.Select(
		"count(*)",
	).From("timekeys")

	k := s.getWhereKeysByFilter(f)

	submissionCount = submissionCount.Where(k)

	// Build the SQL statement and the accomponing arguments.
	sql, args, err := submissionCount.ToSql()

	err = s.db.QueryRowContext(ctx, sql, args...).Scan(&count)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("failed query row context with k: %v", k)
	}
	return count, err
}
