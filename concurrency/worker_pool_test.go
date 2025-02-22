package concurrency

import (
	"reflect"
	"testing"
)

func TestProcessJobs(t *testing.T) {
	jobs := []int{1, 2, 3, 4, 5}
	expected := []int{2, 4, 6, 8, 10} // Each job value is doubled

	result := ProcessJobs(jobs, 3)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
