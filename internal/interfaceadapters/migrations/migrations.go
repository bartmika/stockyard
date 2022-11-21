package migrations

import (
	"database/sql"
	"embed"

	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog/log"
)

//go:embed sql/*.sql
var embedMigrations embed.FS

func RunOnDB(db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Error().Err(err).Caller()
		return err
	}

	// We want to check the version of the database before we run the goose `up`
	// command so we'll know if are starting (version=0) or already started
	// (version!=0). Knowing whether version is 0 or not is valuable for
	// running initialization script.
	version, err := goose.GetDBVersion(db)
	if err != nil {
		log.Error().Err(err).Caller()
		return err
	}

	if err := goose.Up(db, "sql"); err != nil {
		log.Error().Err(err).Caller()
		return err
	}

	if version == 0 {
		err := runMigrate001()
		if err != nil {
			log.Error().Err(err).Caller()
			return err
		}
	}

	return nil
}

func runMigrate001() error {
	log.Info().Msg("NEW DATABASE DETECTED - POPULATING CONTENT")
	return nil
}
