package repository

import (
	"context"
	"encoding/json"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/observation_analyzer_request"
)

// func (s *ObservationAnalyzerRequestRepoImpl) getWhereKeysByFilter(f *domain.ObservationAnalyzerRequestFilter) sq.And {
// 	// Apply specific 'where' keys to apply.
// 	k := sq.And{}
//
// 	if len(f.EntityIDs) > 0 {
// 		entityIDsKey := sq.Or{}
// 		for _, entityID := range f.EntityIDs {
// 			entityIDsKey = append(entityIDsKey, sq.Eq{"entity_id": entityID})
// 		}
// 		k = append(k, entityIDsKey)
// 	}
//
// 	if !f.TimestampGreaterThenOrEqual.IsZero() {
// 		k = append(k, sq.Or{
// 			sq.Gt{"timestamp": f.TimestampGreaterThenOrEqual},
// 			sq.Eq{"timestamp": f.TimestampGreaterThenOrEqual},
// 		})
// 	}
// 	if !f.TimestampGreaterThen.IsZero() {
// 		k = append(k, sq.Gt{"timestamp": f.TimestampGreaterThen})
// 	}
// 	if !f.TimestampLessThen.IsZero() {
// 		k = append(k, sq.Lt{"timestamp": f.TimestampLessThen})
// 	}
// 	if !f.TimestampLessThenOrEqual.IsZero() {
// 		k = append(k, sq.Or{
// 			sq.Lt{"timestamp": f.TimestampLessThenOrEqual},
// 			sq.Eq{"timestamp": f.TimestampLessThenOrEqual},
// 		})
// 	}
//
// 	return k
// }

func (s *ObservationAnalyzerRequestRepoImpl) ListAll(ctx context.Context) ([]*domain.ObservationAnalyzerRequest, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)
	rds := psql.Select(
		"entity_id",
		"uuid",
		"timestamp",
		"type",
		"observation",
	).From("observation_analyzer_requests")

	rds = rds.OrderBy("uuid" + " " + "ASC")

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
		s.logger.Error().Err(err).Caller().Msg("failed prepare")
		return nil, err
	}

	rows, err := stmt.QueryContext(ctx, args...)
	if err != nil {
		s.logger.Error().Err(err).Caller().Msg("failed query context")
		return nil, err
	}

	var arr []*domain.ObservationAnalyzerRequest
	defer rows.Close()
	for rows.Next() {
		var obin []byte
		m := new(domain.ObservationAnalyzerRequest)
		err := rows.Scan(
			&m.EntityID,
			&m.UUID,
			&m.Timestamp,
			&m.Type,
			&obin,
		)
		if err != nil {
			s.logger.Error().Err(err).Caller().Msg("database scan error")
			return nil, err
		}
		if err := json.Unmarshal(obin, &m.Observation); err != nil {
			return nil, err
		}
		arr = append(arr, m)
	}
	err = rows.Err()
	if err != nil {
		s.logger.Error().Err(err).Caller().Msgf("database")
		return nil, err
	}

	if arr == nil {
		return []*domain.ObservationAnalyzerRequest{}, nil
	}
	return arr, err
}
