package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_summation"
)

func (s *ObservationSummationRepoImpl) getWhereKeysByFilter(f *domain.ObservationSummationFilter) sq.And {
	// Apply specific 'where' keys to apply.
	k := sq.And{}

	if len(f.EntityIDs) > 0 {
		entityIDsKey := sq.Or{}
		for _, entityID := range f.EntityIDs {
			entityIDsKey = append(entityIDsKey, sq.Eq{"entity_id": entityID})
		}
		k = append(k, entityIDsKey)
	}

	if !f.StartGreaterThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Gt{"start": f.StartGreaterThenOrEqual},
			sq.Eq{"start": f.StartGreaterThenOrEqual},
		})
	}
	if !f.StartGreaterThen.IsZero() {
		k = append(k, sq.Gt{"start": f.StartGreaterThen})
	}
	if !f.FinishLessThen.IsZero() {
		k = append(k, sq.Lt{"finish": f.FinishLessThen})
	}
	if !f.FinishLessThenOrEqual.IsZero() {
		k = append(k, sq.Or{
			sq.Lt{"finish": f.FinishLessThenOrEqual},
			sq.Eq{"finish": f.FinishLessThenOrEqual},
		})
	}

	return k
}

func (s *ObservationSummationRepoImpl) ListByFilter(ctx context.Context, f *domain.ObservationSummationFilter) ([]*domain.ObservationSummation, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
		"entity_id",
		"start",
		"finish",
		"day",
		"week",
		"month",
		"year",
		"frequency",
		"result",
	).From("observation_summations")
	k := s.getWhereKeysByFilter(f)

	rds = rds.Where(k)

	rds = rds.OrderBy("start" + " " + "ASC")

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

	var arr []*domain.ObservationSummation
	defer rows.Close()
	for rows.Next() {
		m := new(domain.ObservationSummation)
		err := rows.Scan(
			&m.EntityID,
			&m.Start,
			&m.Finish,
			&m.Day,
			&m.Week,
			&m.Month,
			&m.Year,
			&m.Frequency,
			&m.Result,
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
		return []*domain.ObservationSummation{}, nil
	}
	return arr, err
}

func (s *ObservationSummationRepoImpl) CountByFilter(ctx context.Context, f *domain.ObservationSummationFilter) (uint64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// The result we are looking for.
	var count uint64

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	submissionCount := psql.Select(
		"count(*)",
	).From("observation_summations")

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
