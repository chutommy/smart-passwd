package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"syscall/js"
	"time"

	"github.com/chutommy/smart-passwd/pkg/data"
	"github.com/chutommy/smart-passwd/pkg/wasm"
)

var (
	MongoReadURI = "mongodb+srv://appclient:JEZt8LMrj2riAB0C@production.iz3mf.mongodb.net/wordlist?retryWrites=true&w=majority"
	Runner       *wasm.Runner
)

func main() {
	defer os.Exit(0)
	c := make(chan struct{}, 0)

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wl, err := data.ConnectMongo(ctx, MongoReadURI)
	if err != nil {
		fmt.Printf("failed to connect to mongo database: %s\n", err.Error())
		runtime.Goexit()
	}

	Runner = wasm.NewRunner(wl)

	js.Global().Set("gen", js.FuncOf(gen))
	js.Global().Set("stop", js.FuncOf(stop))

	<-c
}

func gen(_ js.Value, p []js.Value) interface{} {
	w, err := Runner.Gen(int16(p[0].Int()))
	if err != nil {
		return fmt.Sprintf("failed to retrieve a random word: %s", err.Error())
	}

	return js.ValueOf(w)
}

func stop(js.Value, []js.Value) interface{} {
	if err := Runner.Stop(); err != nil {
		return fmt.Sprintf("failed to retrieve a random word: %s", err.Error())
	}

	return nil
}