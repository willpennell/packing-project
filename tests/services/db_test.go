package service_test

import (
	"log"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/willpennell/packing-project/service"
)

func TestNewDBService(t *testing.T) {
	dbService := service.NewDBService()

	expectedPacks := []int{5000, 2000, 1000, 500, 250}
	actualPacks, err := dbService.ListAllPackSizes()
	if err != nil {
		log.Fatalf("failed to get pack size: %v", err)
	}
	slices.Sort(expectedPacks)
	slices.Sort(actualPacks)
	assert.Equal(t, expectedPacks, actualPacks)
}

// AddPackSize
func TestAddNewPackSize(t *testing.T) {
	dbService := service.NewDBService()

	size := 15
	err := dbService.AddPackSize(size)
	assert.Nil(t, err)

}

func TestAddExistingPackSize(t *testing.T) {
	dbService := service.NewDBService()

	size := 250
	err := dbService.AddPackSize(size)
	assert.Error(t, err)

}

// RemovePackSize
func TestRemoveExistingPackSize(t *testing.T) {
	dbService := service.NewDBService()

	size := 250
	err := dbService.RemovePackSize(size)
	assert.Nil(t, err)
}

func TestRemoveNonExistingPackSize(t *testing.T) {
	dbService := service.NewDBService()

	size := 15
	err := dbService.RemovePackSize(size)
	assert.Error(t, err)
}

func TestRemoveWhenOnlyOnePackSizeRemains(t *testing.T) {
	dbService := service.NewDBService()
	packsToRemove := []int{5000, 2000, 1000, 500}

	for _, pack := range packsToRemove {
		err := dbService.RemovePackSize(pack)
		assert.Nil(t, err)
	}

	size := 250
	err := dbService.RemovePackSize(size)
	assert.Error(t, err)
}

func TestAddNewPackSizeListAllPacks(t *testing.T) {
	dbService := service.NewDBService()
	expectedResult := []int{5000, 2000, 1000, 500, 250, 15}

	size := 15
	err := dbService.AddPackSize(size)
	assert.Nil(t, err)

	actualPackSizes, err := dbService.ListAllPackSizes()
	assert.Nil(t, err)

	slices.Sort(expectedResult)
	slices.Sort(actualPackSizes)
	assert.Equal(t, expectedResult, actualPackSizes)
}
