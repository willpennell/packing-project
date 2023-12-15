package router

import (
	"github.com/gin-gonic/gin"
	"github.com/willpennell/packing-project/handler"
)

func InitializeRoutes(router *gin.Engine) {
	packRoutes := router.Group("/pack")
	{
		packRoutes.GET("/:number", handler.GetPackHandler)
		packRoutes.POST("/add/:number")
		packRoutes.DELETE("/delete/:number")

		packRoutes.GET("/list")
		packRoutes.POST("/new")
		packRoutes.GET("/reset")
	}
}
