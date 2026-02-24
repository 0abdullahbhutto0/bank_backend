package main

import (
	"context"
	"log"

	"github.com/0abdullahbhutto0/bank_backend/api"
	db "github.com/0abdullahbhutto0/bank_backend/db/sqlc"
	"github.com/0abdullahbhutto0/bank_backend/util"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load config")
	}
	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Cannot Connect to the DB.", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
