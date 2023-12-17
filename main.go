package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/willpennell/packing-project/router"
	"github.com/willpennell/packing-project/service"
)

func main() {
	dbService := service.NewDBService()
	r := gin.Default()
	r.LoadHTMLGlob("templates/**")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	router.InitializeRoutes(r, dbService)
	r.Run()

}
