package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/willpennell/packing-project/router"
	"github.com/willpennell/packing-project/service"
)

func main() {
	items := 12001

	packSizes := []int{5000, 2000, 1000, 500, 250}

	packer := service.PackService{}

	fmt.Println(packer.PackItems(packSizes, items))

	r := gin.Default()
	router.InitializeRoutes(r)
	r.Run()
}
