package data

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/chutified/smart-passwd/pkg/utils"
	_ "modernc.org/sqlite"
)

var (
	testDBFile *utils.File
	testDB     *sql.DB
)

func TestMain(m *testing.M) {
	var err error

	testDBFile = utils.NewFile("test", "wordlist", "db")
	if testDB, err = sql.Open("sqlite", testDBFile.FilePath()); err != nil {
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
