package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

var ctx = context.Background()

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
