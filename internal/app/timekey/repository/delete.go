package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	domain "github.com/bartmika/stockyard/internal/domain/timekey"
)

func (s *TimeKeyRepoImpl) DeleteByFilter(ctx context.Context, f *domain.TimeKeyFilter) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	psql := sq.StatementBuilder.RunWith(s.dbCache).PlaceholderFormat(sq.Dollar)

	k := s.getWhereKeysByFilter(f)

	sql, args, err := psql.
		Delete("timekeys").
		Where(k).
		ToSql()

	stmt, err := s.db.PrepareContext(ctx, sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, args...)
	return err

}
