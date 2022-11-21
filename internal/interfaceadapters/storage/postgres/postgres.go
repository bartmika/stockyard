package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/bartmika/stockyard/internal/config"
)

func ConnectDB(c *config.Conf) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable",
		c.DB.Host,
		c.DB.Port,
		c.DB.User,
		c.DB.Password,
		c.DB.DBName,
	)

	dbInstance, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = dbInstance.Ping()
	if err != nil {
		return nil, err
	}

	return dbInstance, nil
}
