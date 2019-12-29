package sort

import "testing"

// go test -cover coverage command

func TestBubbleSortIncreasingOrder(t *testing.T) {
	// Init
	elements := GetElements(10)

	// Execution
	BubbleSort(elements)
	// Validation
	// ここのケースは全て通す必要がある
	// カバレージはあくまで指針でエンジニア側でそのエビデンスを取る必要がある
	if elements[0] != 10 {
		t.Error("first elements should be 9")
	}

	if elements[len(elements)-1] != 1 {
		t.Error("last element shoul be 0")
	}

}

func TestSortIncreasingOrder(t *testing.T) {
	elements := GetElements(10)

	Sort(elements)

	if elements[0] != 1 {
		t.Error("first elements should be 1")
	}

	if elements[len(elements)-1] != 10 {
		t.Error("last element shoul be 9")
	}

}

func BenchmarkBubbleSort(b *testing.B) {
	elements := GetElements(100000)

	for i := 0; i < b.N; i++ {
		BubbleSort(elements)
	}
}

func BenchmarkSort(b *testing.B) {
	elements := GetElements(10000)

	for i := 0; i < b.N; i++ {
		Sort(elements)
	}
}
