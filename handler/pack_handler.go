package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/willpennell/packing-project/model"
	"github.com/willpennell/packing-project/service"
)

func GetPackHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		numberStr := c.Param("number")
		number, err := strconv.Atoi(numberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number"})
			return
		}

		packSizes, err := dbService.ListAllPackSizes()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		packer := service.PackService{}
		response, err := packer.PackItems(packSizes, number)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, response)
	}
}

func AddPackSizeHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		numberStr := c.Query("size")
		if numberStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query"})
			return
		}
		packSize, err := strconv.Atoi(numberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
			return
		}
		if err := dbService.AddPackSize(packSize); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Status(http.StatusOK)
	}
}

func RemovePackSizeHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		numberStr := c.Query("size")
		if numberStr == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid query"})
			return
		}
		packSize, err := strconv.Atoi(numberStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid number"})
			return
		}
		if err := dbService.RemovePackSize(packSize); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		c.Status(http.StatusOK)
	}
}

func ListAllPackSizesHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		packSizes, err := dbService.ListAllPackSizes()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"pack_sizes": packSizes})
	}
}

func CreateNewListOfPackSizesHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.NewPacksRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		newPackSizes := req.PackSizes

		packer := service.PackService{}
		newPackSizes = packer.RemoveDuplicates(newPackSizes)

		if len(newPackSizes) < 1 {
			err := errors.New("not enough pack sizes")
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		if err := dbService.ResetPackSize(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		if err := dbService.NewPackSizes(newPackSizes); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.Status(http.StatusOK)
	}
}

func ResetPackSizesHandler(dbService *service.DBService) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := dbService.ResetPackSize(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.Status(http.StatusOK)
	}
}
