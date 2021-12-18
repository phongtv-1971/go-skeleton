package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/phongtv-1971/go-skeleton/util"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..", "test")
	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
