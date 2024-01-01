package server

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/do"
)

func Init(
	router gin.IRouter,
	injector *do.Injector,
) error {
	service := NewService()

	router.GET("/health", service.Health)

	return nil
}
