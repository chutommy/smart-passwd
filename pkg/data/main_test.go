package data

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	if testDB, err = sql.Open("sqlite3", "test/wordlist.db"); err != nil {
		log.Fatal(err)
	}

	if testDB == nil {
		log.Fatal("nil value of testDB")
	}

	if err = testDB.Ping(); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
