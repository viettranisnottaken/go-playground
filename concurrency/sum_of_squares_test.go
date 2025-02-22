package concurrency

import "testing"

func TestSumOfSquares(t *testing.T) {

	nums := make([]int, 1000000)

	for i := 0; i < 1000000; i++ {
		nums[i] = i + 1
	}

	expected := 333333833333500000

	result := SumOfSquares(nums)
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}
