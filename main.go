package main

import (
	"log"
	"wblitz-watcher/app"
)

func main() {
	err := app.New().Serve()
	if err != nil {
		log.Fatalf("%+v", err)
	}
}
