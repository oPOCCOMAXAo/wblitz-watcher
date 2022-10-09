package app

import (
	"context"
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
	ctx, cancelFn := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancelFn()

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

	go app.TaskClans(ctx)

	<-ctx.Done()

	return app.Close()
}
