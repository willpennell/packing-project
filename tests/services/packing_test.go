package service_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/willpennell/packing-project/service"

	"github.com/stretchr/testify/assert"
)

func TestPackItems(t *testing.T) {
	packer := service.PackService{}

	packSizes := []int{10, 5, 2}
	items := 15

	expectedExtraItems := 0
	expectedTotalItems := 15
	expectedTotalPacks := 2

	res, err := packer.PackItems(packSizes, items)
	if err != nil {
		log.Fatalf(err.Error())
	}

	assert.Equal(t, expectedExtraItems, res.ExtraItems)
	assert.NotNil(t, res.Packs)
	assert.Equal(t, expectedTotalItems, res.TotalItems)
	assert.Equal(t, expectedTotalPacks, res.TotalPacks)
}

func TestEmptyList(t *testing.T) {
	packer := service.PackService{}

	packSizes := []int{}
	items := 10

	fmt.Print(len(packSizes))
	_, err := packer.PackItems(packSizes, items)
	assert.Error(t, err)
}

func TestDuplicatePackSizes(t *testing.T) {
	packer := service.PackService{}

	packSizes := []int{10, 10, 5}
	items := 10

	expectedExtraItems := 0
	expectedTotalItems := 10
	expectedTotalPacks := 1
	expectedPackCounts := []int{1, 0}

	packer.PackItems(packSizes, items)

	res, err := packer.PackItems(packSizes, items)
	if err != nil {
		log.Fatalf(err.Error())
	}

	assert.Equal(t, expectedExtraItems, res.ExtraItems)
	assert.NotNil(t, expectedPackCounts, res.Packs)
	assert.Equal(t, expectedTotalItems, res.TotalItems)
	assert.Equal(t, expectedTotalPacks, res.TotalPacks)

}

func TestRequirementInput(t *testing.T) {
	packer := service.PackService{}
	packSizes := []int{5000, 2000, 1000, 500, 250}
	testScenarios := []struct {
		packSizes          []int
		items              int
		expectedExtraItems int
		expectedTotalItems int
		expectedTotalPacks int
		expectedPackCounts []int
	}{
		{packSizes, 1, 249, 250, 1, []int{0, 0, 0, 0, 1}},
		{packSizes, 250, 0, 250, 1, []int{0, 0, 0, 0, 1}},
		{packSizes, 251, 249, 500, 1, []int{0, 0, 0, 1, 0}},
		{packSizes, 501, 249, 750, 2, []int{0, 0, 0, 1, 1}},
		{packSizes, 12001, 249, 12250, 4, []int{2, 1, 0, 0, 1}},
	}

	for _, tt := range testScenarios {
		res, err := packer.PackItems(tt.packSizes, tt.items)
		if err != nil {
			t.Fatalf(err.Error())
		}

		assert.Equal(t, tt.expectedExtraItems, res.ExtraItems)
		assert.NotNil(t, tt.expectedPackCounts, res.Packs)
		assert.Equal(t, tt.expectedTotalItems, res.TotalItems)
		assert.Equal(t, tt.expectedTotalPacks, res.TotalPacks)
	}
}

func TestZeroItems(t *testing.T) {
	packer := service.PackService{}
	packSizes := []int{5000, 2000, 1000, 500, 250}
	items := 0

	_, err := packer.PackItems(packSizes, items)
	assert.Error(t, err)
}
