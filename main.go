package main

import (
	"github.com/gin-gonic/gin"

	"github.com/willpennell/packing-project/router"
	"github.com/willpennell/packing-project/service"
)

func main() {
	dbService := service.NewDBService()
	r := gin.Default()
	router.InitializeRoutes(r, dbService)
	r.Run()

}
