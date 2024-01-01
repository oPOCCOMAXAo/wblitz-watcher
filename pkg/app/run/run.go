package run

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/app/config"
	"github.com/opoccomaxao/wblitz-watcher/pkg/endpoints"
	"github.com/opoccomaxao/wblitz-watcher/pkg/server"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/services/watcher"
)

//nolint:funlen
func Run() error {
	appCtx, cancelFn := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancelFn()

	appCtx, cancelCauseFn := context.WithCancelCause(appCtx)
	defer cancelCauseFn(nil)

	config, err := config.Load()
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	injector := do.New()
	defer func() {
		err := injector.Shutdown()
		if err != nil {
			log.Printf("%+v\n", err)
		}
	}()

	config.Provide(injector, appCtx)

	server, err := server.Invoke(injector)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = endpoints.Init(server.Router(), injector)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	discord, err := discord.Invoke(injector)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = discord.Serve()
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	watcher, err := watcher.Invoke(injector)
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	err = watcher.Serve()
	if err != nil {
		//nolint:wrapcheck
		return err
	}

	go func() {
		cancelCauseFn(server.Serve())
	}()

	<-appCtx.Done()

	err = appCtx.Err()
	if err != nil {
		log.Printf("%+v\n", err)
	}

	return nil
}
