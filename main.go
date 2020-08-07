package main

import (
	"fmt"
	"log"
	"time"

	controls "github.com/chutified/smart-passwd/controls"
	_ "github.com/lib/pq"
)

func main() {
	t0 := time.Now()

	c := controls.New()
	err := c.Init()
	if err != nil {
		log.Fatal(err)
	}
	defer c.Stop()

	err = c.Gen(5)
	if err != nil {
		fmt.Printf("err = %+v\n", err)
	}

	fmt.Println(time.Since(t0))
}
