package engine

import (
	"log"
	"os"
	"testing"

	"github.com/chutified/smart-passwd/pkg/data"
	"github.com/chutified/smart-passwd/pkg/utils"
	_ "github.com/mattn/go-sqlite3"
)

var (
	testEngine        *Engine
	testInvalidEngine *Engine
)

func TestMain(m *testing.M) {
	wl, err := data.Connect(utils.NewFile("test", "wordlist", "db"))
	if err != nil {
		log.Fatal(err)
	}

	c := NewConstructor(3, 20)
	s := NewSwapper(Alphabet(), Specials(), SwapList())

	testEngine = Init(wl, c, s)
	testInvalidEngine = Init(&data.WordList{}, c, s)

	os.Exit(m.Run())
}