package service

import (
	"errors"
	"fmt"
	"sort"
)

type Packer interface {
	PackItems(packSizes []int, items int) (int, int, int, []int)
	RemoveDuplicates(packSizes []int) []int
}

type PackService struct{}

func (ps PackService) PackItems(packSizes []int, items int) (int, int, int, []int, error) {

	if items <= 0 {
		return 0, 0, 0, nil, errors.New("you need more than 0 items for an order")
	}
	if len(packSizes) <= 0 {
		return 0, 0, 0, nil, errors.New("you must have pack sizes to send")
	}

	packSizes = ps.RemoveDuplicates(packSizes)
	sortSizesDescending(packSizes)

	counters := make([]int, len(packSizes))
	fmt.Println(packSizes)
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

	// return struct
	return extraItems, totalItems, totalPacks, counters, nil
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
