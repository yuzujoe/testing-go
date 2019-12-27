package services

import "testing-go/api/utils/sort"

const (
	privateConst = "private"
)

func Sort(elements []int) {
	sort.BubbleSort(elements)
}
