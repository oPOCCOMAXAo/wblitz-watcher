package discord

import "github.com/gin-gonic/gin"

func (s *Service) Interactions(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"type": 1,
	})
}
