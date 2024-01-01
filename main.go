package main

import (
	"log"

	"github.com/opoccomaxao/wblitz-watcher/pkg/app/run"
)

func main() {
	err := run.Run()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
