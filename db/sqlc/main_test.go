package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/0abdullahbhutto0/bank_backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	config, erro := util.LoadConfig("../..")
	if erro != nil {
		log.Fatal("Could not load config", erro)
	}

	var err error
	testDB, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to the DB.", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
