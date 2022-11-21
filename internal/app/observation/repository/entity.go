package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

type ObservationRepoImpl struct {
	logger  *zerolog.Logger
	db      *sql.DB
	dbCache *sq.StmtCache
}

func NewObservationRepoImpl(db *sql.DB, logger *zerolog.Logger) *ObservationRepoImpl {
	return &ObservationRepoImpl{
		logger:  logger,
		db:      db,
		dbCache: sq.NewStmtCache(db),
	}
}
