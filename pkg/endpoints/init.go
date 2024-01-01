package endpoints

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"

	"github.com/opoccomaxao/wblitz-watcher/pkg/endpoints/discord"
	"github.com/opoccomaxao/wblitz-watcher/pkg/endpoints/server"
)

type Endpoints = func(gin.IRouter, *do.Injector) error

func Init(
	router gin.IRouter,
	injector *do.Injector,
) error {
	return finish(
		router,
		injector,
		server.Init,
		discord.Init,
	)
}

func finish(
	router gin.IRouter,
	injector *do.Injector,
	endpoints ...Endpoints,
) error {
	for _, init := range endpoints {
		err := init(router, injector)
		if err != nil {
			return err
		}
	}

	return nil
}
