package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/willpennell/packing-project/services"
)

func main() {
	items := 12001

	packSizes := []int{5000, 2000, 1000, 500, 250}

	packer := services.PackService{}

	fmt.Println(packer.PackItems(packSizes, items))

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello World, its wills api")
	})
	r.Run()
}
