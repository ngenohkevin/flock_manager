package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ngenohkevin/flock_manager/api"
	db "github.com/ngenohkevin/flock_manager/db/sqlc"
	"log"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:85dilanwest@localhost:5432/flock_manager?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("unable to start the server", err)
	}
}
