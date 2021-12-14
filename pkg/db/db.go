package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres drivers

	"github.com/abdullohsattorov/todo-service/config"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	psqlString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)
	connDb, err := sqlx.Connect("postgres", psqlString)
	if err != nil {
		return nil, err
	}
	return connDb, nil
}
