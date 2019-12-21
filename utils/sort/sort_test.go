package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	// Init
	elements := []int{9, 8, 7, 5, 3, 2, 1, 4, 6, 0}

	// Execution
	BubbleSort(elements)

	// Validation
	if elements[0] != 9 {
		t.Error("first elements should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element shoul be 0")
	}

}
