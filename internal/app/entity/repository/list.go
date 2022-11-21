package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/entity"
)

func (s *EntityRepoImpl) getWhereKeysByFilter(f *domain.EntityFilter) sq.And {
	// Apply specific 'where' keys to apply.
	k := sq.And{}

	// Here is where you would add filter keys...
	if f.DataType != 0 {
		k = append(k, sq.Eq{"data_type": f.DataType})
	}

	// if len(f.States) > 0 {
	// 	statesKey := sq.Or{}
	// 	for _, state := range f.States {
	// 		statesKey = append(statesKey, sq.Eq{"state": state})
	// 	}
	// 	k = append(k, statesKey)
	// }

	if len(f.IDs) > 0 {
		pksKey := sq.Or{}
		for _, pk := range f.IDs {
			pksKey = append(pksKey, sq.Eq{"id": pk})
		}
		k = append(k, pksKey)
	}

	return k
}

func (s *EntityRepoImpl) ListByFilter(ctx context.Context, f *domain.EntityFilter) ([]*domain.Entity, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
		"id",
		"uuid",
		"name",
		"data_type",
	).From("entities")
	k := s.getWhereKeysByFilter(f)

	rds = rds.Where(k)

	rds = rds.OrderBy(f.SortField + " " + f.SortOrder)

	// Note:
	// (1) https://ivopereira.net/efficient-pagination-dont-use-offset-limit
	// (2) https://github.com/Masterminds/squirrel/blob/def598cbb358368fbfc3f6a9a914699a36846992/select_test.go#L41

	rds = rds.Offset(f.Offset).Suffix("FETCH FIRST ? ROWS ONLY", f.Limit)

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

	var arr []*domain.Entity
	defer rows.Close()
	for rows.Next() {
		m := new(domain.Entity)
		err := rows.Scan(
			&m.ID,
			&m.UUID,
			&m.Name,
			&m.DataType,
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
		return []*domain.Entity{}, nil
	}
	return arr, err
}

func (s *EntityRepoImpl) CountByFilter(ctx context.Context, f *domain.EntityFilter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// The result we are looking for.
	var count uint64

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	submissionCount := psql.Select(
		"count(*)",
	).From("entities")

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
