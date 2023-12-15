package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/willpennell/packing-project/service"
)

func GetPackHandler(c *gin.Context) {
	numberStr := c.Param("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
		return
	}

	packer := service.PackService{}
	packSizes := []int{5000, 2000, 1000, 500, 250}
	extraItems, totalItems, totalPacks, counts, err := packer.PackItems(packSizes, number)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"extraItems": extraItems, "totalItems": totalItems, "totalPacks": totalPacks, "counts": counts})
}
