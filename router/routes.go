package router

import (
	"github.com/gin-gonic/gin"
	"github.com/willpennell/packing-project/handler"
	"github.com/willpennell/packing-project/service"
)

func InitializeRoutes(router *gin.Engine, dbService *service.DBService) {
	packRoutes := router.Group("/pack")
	{
		packRoutes.GET("/:number", handler.GetPackHandler(dbService))
		packRoutes.POST("/add", handler.AddPackSizeHandler(dbService))
		packRoutes.DELETE("/delete", handler.RemovePackSizeHandler(dbService))
		packRoutes.GET("/list", handler.ListAllPackSizesHandler(dbService))
		packRoutes.POST("/new", handler.CreateNewListOfPackSizesHandler(dbService))
		packRoutes.GET("/reset", handler.ResetPackSizesHandler(dbService))
	}
}
