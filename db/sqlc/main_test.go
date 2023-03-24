package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/ngenohkevin/flock_manager/db/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
