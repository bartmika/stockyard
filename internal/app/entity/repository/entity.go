package repository

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	"github.com/rs/zerolog"
)

type EntityRepoImpl struct {
	logger  *zerolog.Logger
	db      *sql.DB
	dbCache *sq.StmtCache
}

func NewEntityRepoImpl(db *sql.DB, logger *zerolog.Logger) *EntityRepoImpl {
	return &EntityRepoImpl{
		logger:  logger,
		db:      db,
		dbCache: sq.NewStmtCache(db),
	}
}
