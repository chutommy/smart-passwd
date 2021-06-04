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
	testDBFile      *utils.File
	testSQLWordList *SQLiteWordList
)

func TestMain(m *testing.M) {
	var err error

	testDBFile = utils.NewFile("test", "wordlist", "db")
	testSQLWordList = &SQLiteWordList{}

	if testSQLWordList.db, err = sql.Open("sqlite", testDBFile.FilePath()); err != nil {
		log.Fatal(err)
	}

	if testSQLWordList.db == nil {
		log.Fatal("nil value of testDB")
	}

	if err = testSQLWordList.db.Ping(); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
