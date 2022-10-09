package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"wblitz-watcher/repo"
	"wblitz-watcher/sender"
	"wblitz-watcher/wg/client"

	"github.com/opoccomaxao-go/task-server/task"
)

type App struct {
	storage task.Storage
	client  *client.Client
	repo    *repo.Repo
	sender  *sender.Client
	config  Config
}

func New() *App {
	return &App{}
}

func (app *App) Serve() error {
	closeSignal := make(chan os.Signal, 10)
	signal.Notify(closeSignal, syscall.SIGINT, syscall.SIGTERM)

	err := app.initConfig()
	if err != nil {
		return err
	}

	err = app.initClient()
	if err != nil {
		return err
	}

	err = app.initSender()
	if err != nil {
		return err
	}

	err = app.initRepo()
	if err != nil {
		return err
	}

	err = app.ProcessClans(context.Background())
	if err != nil {
		return err
	}

	// <-closeSignal

	return app.Close()
}
