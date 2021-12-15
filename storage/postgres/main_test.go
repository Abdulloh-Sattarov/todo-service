package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/abdullohsattorov/todo-service/config"
	"github.com/abdullohsattorov/todo-service/pkg/db"
	"github.com/abdullohsattorov/todo-service/pkg/logger"
)

var pgRepo *todoRepo

func TestMain(m *testing.M) {
	cfg := config.Load()

	connDB, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	pgRepo = NewTodoRepo(connDB)

	os.Exit(m.Run())
}
