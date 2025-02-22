package concurrency

import "testing"

func TestFindMax(t *testing.T) {
	nums := []int{3, 7, 2, 8, 5, 9, 1}
	expected := 9

	result := FindMax(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
