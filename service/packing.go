package service

import (
	"errors"
	"fmt"
	"sort"

	"github.com/willpennell/packing-project/model"
)

type Packer interface {
	PackItems(packSizes []int, items int) (model.PackedItemsInfo, error)
	RemoveDuplicates(packSizes []int) []int
}

type PackService struct{}

func (ps PackService) PackItems(packSizes []int, items int) (model.PackedItemsInfo, error) {

	if items <= 0 {
		return model.PackedItemsInfo{}, errors.New("you need more than 0 items for an order")
	}
	if len(packSizes) <= 0 {
		return model.PackedItemsInfo{}, errors.New("you must have pack sizes to send")
	}

	packSizes = ps.RemoveDuplicates(packSizes)
	sortSizesDescending(packSizes)

	counters := make([]int, len(packSizes))
	minValue := packSizes[len(packSizes)-1]

	for i, pack := range packSizes {
		count := items / pack
		items -= count * pack
		counters[i] = count
		if items == 0 {
			break
		}
	}

	extraItems := 0
	if items > 0 && items < minValue {
		counters[len(counters)-1] += 1
		extraItems = minValue - items
		items = 0
	}

	resizeToLargerPacks(packSizes, counters)
	totalItems, totalPacks := totals(packSizes, counters)

	packs := make(map[string]model.PackInfo)
	for i, size := range packSizes {
		key := fmt.Sprintf("box_%d", i+1)
		packs[key] = model.PackInfo{
			Size: size,
			Used: counters[i],
		}
	}

	response := model.PackedItemsInfo{
		ExtraItems: extraItems,
		TotalItems: totalItems,
		TotalPacks: totalPacks,
		Packs:      packs,
	}
	return response, nil
}

func resizeToLargerPacks(packSizes []int, counters []int) {
	for i := len(counters) - 1; i > 0; i-- {
		combinedValue := packSizes[i] * counters[i]
		if combinedValue >= packSizes[i-1] {
			packsToCombine := combinedValue / packSizes[i-1]
			packsToRemove := packsToCombine * (packSizes[i-1] / packSizes[i])
			counters[i] -= packsToRemove
			counters[i-1] += packsToCombine
		}
	}
}

func totals(packSizes []int, counters []int) (int, int) {
	totalItemsSent := 0
	totalPacks := 0
	for i, count := range counters {
		totalItemsSent += count * packSizes[i]
		totalPacks += count
	}
	return totalItemsSent, totalPacks
}

func sortSizesDescending(packSizes []int) {
	sort.Slice(packSizes, func(i, j int) bool {
		return packSizes[i] > packSizes[j]
	})
}

func (ps PackService) RemoveDuplicates(listToRemove []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range listToRemove {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
