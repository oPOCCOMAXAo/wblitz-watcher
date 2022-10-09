package main

import (
	"log"

	"github.com/opoccomaxao/wblitz-watcher/app"
)

func main() {
	err := app.New().Serve()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
