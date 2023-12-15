package services_test

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
	expectedPackCounts := []int{1, 1, 0}

	extraItems, actualTotalItems, actualTotalPack, actualCounts, err := packer.PackItems(packSizes, items)
	if err != nil {
		log.Fatalf(err.Error())
	}

	assert.Equal(t, expectedExtraItems, extraItems)
	assert.Equal(t, expectedPackCounts, actualCounts)
	assert.Equal(t, expectedTotalItems, actualTotalItems)
	assert.Equal(t, expectedTotalPacks, actualTotalPack)
}

func TestEmptyList(t *testing.T) {
	packer := service.PackService{}

	packSizes := []int{}
	items := 10

	fmt.Print(len(packSizes))
	_, _, _, _, err := packer.PackItems(packSizes, items)
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

	extraItems, actualTotalItems, actualTotalPack, actualCounts, err := packer.PackItems(packSizes, items)
	if err != nil {
		log.Fatalf(err.Error())
	}

	assert.Equal(t, expectedExtraItems, extraItems)
	assert.Equal(t, expectedPackCounts, actualCounts)
	assert.Equal(t, expectedTotalItems, actualTotalItems)
	assert.Equal(t, expectedTotalPacks, actualTotalPack)

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
		extraItems, actualTotalItems, actualTotalPack, actualCounts, err := packer.PackItems(tt.packSizes, tt.items)
		if err != nil {
			t.Fatalf(err.Error())
		}

		assert.Equal(t, tt.expectedExtraItems, extraItems)
		assert.Equal(t, tt.expectedPackCounts, actualCounts)
		assert.Equal(t, tt.expectedTotalItems, actualTotalItems)
		assert.Equal(t, tt.expectedTotalPacks, actualTotalPack)
	}
}

func TestZeroItems(t *testing.T) {
	packer := service.PackService{}
	packSizes := []int{5000, 2000, 1000, 500, 250}
	items := 0

	_, _, _, _, err := packer.PackItems(packSizes, items)
	assert.Error(t, err)
}
