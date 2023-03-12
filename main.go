package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/lushenle/simplebank/api"
	db "github.com/lushenle/simplebank/db/sqlc"
)

const (
	driverName     = "postgres"
	dataSourceName = "postgresql://myuser:mypass@localhost:5432/simplebank?sslmode=disable"
	serverAdderess = ":8080"
)

func main() {
	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAdderess)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
