package services

import "testing"

func TestConstants(t *testing.T) {
	if privateConst != "private" {
		t.Error("privateConst should be private")
	}
}
func TestSort(t *testing.T) {
	elements := []int{7, 9, 5, 3, 1, 0, 4, 6, 8, 2}

	Sort(elements)

	if elements[0] != 9 {
		t.Error("first elements should be 9")
	}

	if elements[len(elements)-1] != 0 {
		t.Error("last element shoul be 0")
	}
}
