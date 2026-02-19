package main

import (
	"context"
	"log"

	"github.com/0abdullahbhutto0/bank_backend/api"
	db "github.com/0abdullahbhutto0/bank_backend/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	//dbDriver = "postgres"
	dbSource      = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress = "0.0.0.0:8050"
)

func main() {
	conn, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("Cannot Connect to the DB.", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}
