package sort

import "testing"

// go test -cover coverage command

func TestBubbleSortOrderDESC(t *testing.T) {
	// Init
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}

	// Execution
	BubbleSort(elements)
	// Validation
	// ここのケースは全て通す必要がある
	if elements[0] != 9 {
		t.Error("first elements should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element shoul be 0")
	}

}
