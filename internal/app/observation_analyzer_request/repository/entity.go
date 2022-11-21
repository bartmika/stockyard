package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

type ObservationAnalyzerRequestRepoImpl struct {
	logger  *zerolog.Logger
	db      *sql.DB
	dbCache *sq.StmtCache
}

func NewObservationAnalyzerRequestRepoImpl(db *sql.DB, logger *zerolog.Logger) *ObservationAnalyzerRequestRepoImpl {
	return &ObservationAnalyzerRequestRepoImpl{
		logger:  logger,
		db:      db,
		dbCache: sq.NewStmtCache(db),
	}
}
