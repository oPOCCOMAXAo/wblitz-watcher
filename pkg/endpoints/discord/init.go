package discord

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Init(
	router gin.IRouter,
	injector *do.Injector,
) error {
	service := NewService()

	router.POST("/interactions", service.Interactions)

	return nil
}
