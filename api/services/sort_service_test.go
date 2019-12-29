package services

import (
	"fmt"
	"testing"
	"testing-go/api/utils/sort"
)

func TestSort(t *testing.T) {
	elements := sort.GetElements(10)
	Sort(elements)

	if elements[0] != 0 {
		t.Error("first elements should be 0")
	}

	if elements[len(elements)-1] != 9 {
		t.Error("last element shoul be 9")
	}
}

func TestSortMoreThan10000(t *testing.T) {
	elements := sort.GetElements(10001)
	Sort(elements)
	fmt.Println(elements)
	if elements[0] != 1 {
		t.Error("first elements should be 0")
	}

	if elements[len(elements)-1] != 10001 {
		t.Error("last element shoul be 10000")
	}
}

func BenchmarkBubbleSort10k(b *testing.B) {
	elements := sort.GetElements(20000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}

func BenchmarkBubbleSort100k(b *testing.B) {
	elements := sort.GetElements(100000)
	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
