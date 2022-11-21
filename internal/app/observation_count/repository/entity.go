package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

type ObservationCountRepoImpl struct {
	logger  *zerolog.Logger
	db      *sql.DB
	dbCache *sq.StmtCache
}

func NewObservationCountRepoImpl(db *sql.DB, logger *zerolog.Logger) *ObservationCountRepoImpl {
	return &ObservationCountRepoImpl{
		logger:  logger,
		db:      db,
		dbCache: sq.NewStmtCache(db),
	}
}
