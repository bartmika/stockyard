package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

type TimeKeyRepoImpl struct {
	logger  *zerolog.Logger
	db      *sql.DB
	dbCache *sq.StmtCache
}

func NewTimeKeyRepoImpl(db *sql.DB, logger *zerolog.Logger) *TimeKeyRepoImpl {
	return &TimeKeyRepoImpl{
		logger:  logger,
		db:      db,
		dbCache: sq.NewStmtCache(db),
	}
}
