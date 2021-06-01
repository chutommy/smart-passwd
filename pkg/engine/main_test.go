package engine

import (
	"log"
	"os"
	"testing"

	"github.com/chutified/smart-passwd/pkg/data"
	"github.com/chutified/smart-passwd/pkg/utils"
	_ "modernc.org/sqlite"
)

var (
	testEngine        *Engine
	testInvalidEngine *Engine
)

func TestMain(m *testing.M) {
	wl, err := data.ConnectSQLite(utils.NewFile("test", "wordlist", "db"))
	if err != nil {
		log.Fatal(err)
	}

	c := NewConstructor(3, 20)
	s := NewSwapper()

	testEngine = Init(wl, c, s)
	testInvalidEngine = Init(&data.SQLiteWordList{}, c, s)

	os.Exit(m.Run())
}
