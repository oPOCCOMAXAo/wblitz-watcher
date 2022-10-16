package app

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/opoccomaxao/wblitz-watcher/repo"
	"github.com/opoccomaxao/wblitz-watcher/sender"
	"github.com/opoccomaxao/wblitz-watcher/wg/client"

	"github.com/opoccomaxao-go/task-server/task"
)

type App struct {
	storage task.Storage
	client  *client.Client
	repo    *repo.Repo
	sender  *sender.Client
	server  *http.Server
	config  Config
}

func New() *App {
	return &App{}
}

func (app *App) Serve() error {
	termSignals := []os.Signal{syscall.SIGINT, syscall.SIGTERM}

	sigs := make(chan os.Signal, 10)
	signal.Notify(sigs, termSignals...)

	ctx, cancelFn := signal.NotifyContext(context.Background(), termSignals...)
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

	err = app.initServer()
	if err != nil {
		return err
	}

	go app.TaskClans(ctx)

	sig := <-sigs
	log.Printf("Received signal: %s\n", sig.String())
	cancelFn()

	return app.Close()
}
