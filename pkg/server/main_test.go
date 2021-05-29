package server

import (
	"log"
	"os"
	"testing"

	"github.com/chutified/smart-passwd/pkg/config"
	"github.com/chutified/smart-passwd/pkg/data"
	"github.com/chutified/smart-passwd/pkg/engine"
	"github.com/chutified/smart-passwd/pkg/utils"
	"github.com/gin-gonic/gin"
	_ "modernc.org/sqlite"
)

var (
	testConfig    *config.Config
	testEngine    *engine.Engine
	invalidEngine *engine.Engine
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	wl, err := data.Connect(utils.NewFile("test", "wordlist", "db"))
	if err != nil {
		log.Fatal(err)
	}

	iwl, err := data.Connect(utils.NewFile("test", "invalid-wordlist", "db"))
	if err != nil {
		log.Fatal(err)
	}

	c := engine.NewConstructor(3, 20)
	s := engine.NewSwapper()

	testConfig = config.NewConfig(8080, "./test/wordlist.db", false, "../..")
	testEngine = engine.Init(wl, c, s)
	invalidEngine = engine.Init(iwl, c, s)

	os.Exit(m.Run())
}
