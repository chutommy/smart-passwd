package main

import (
	"context"
	"log"
	"os"
	"syscall/js"

	engine "github.com/chutommy/smart-passwd/pkg/enginelite"
)

func main() {
	defer os.Exit(0)

	js.Global().Set("distribute", js.FuncOf(distribute))
	js.Global().Set("generate", js.FuncOf(generate))

	c := make(chan struct{}, 0)
	<-c
}

func generate(_ js.Value, p []js.Value) interface{} {
	engi := engine.Init(engine.NewConstructor(3, 22), engine.NewSwapper())

	resp, err := engi.Generate(context.Background(), engine.NewRequest(0, int16(p[0].Int()), p[1].String()))
	if err != nil {
		log.Fatal(err)
	}

	out := make([]interface{}, 2)
	out[0] = resp.Password()
	out[1] = resp.Helper()

	return js.ValueOf(out)
}

func distribute(_ js.Value, p []js.Value) interface{} {
	ctr := engine.NewConstructor(3, 22)

	arr, err := ctr.Distribute(int16(p[0].Int()))
	if err != nil {
		log.Fatal(err)
	}

	out := make([]interface{}, len(arr))
	for i, e := range arr {
		out[i] = e
	}

	return js.ValueOf(out)
}
