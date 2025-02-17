package binary_search

import "testing"

type FindMinInRotatedArrayTestInput struct {
	nums []int
}

func TestFindMinimumInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		inputs FindMinInRotatedArrayTestInput
		expect int
	}{
		{
			name: "Case 1",
			inputs: FindMinInRotatedArrayTestInput{
				nums: []int{3, 4, 5, 6, 1, 2},
			},
			expect: 1,
		},
		{
			name: "Case 2: min is in the middle",
			inputs: FindMinInRotatedArrayTestInput{
				nums: []int{4, 5, 6, 0, 1, 2, 3},
			},
			expect: 0,
		},
		{
			name: "Case 3",
			inputs: FindMinInRotatedArrayTestInput{
				nums: []int{4, 5, 6, 7},
			},
			expect: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := FindMinimumInRotatedSortedArray(tt.inputs.nums)

			if got != tt.expect {
				t.Errorf("FindMinimumInRotatedSortedArray(%v) = %d; expected %d", tt.inputs.nums, got, tt.expect)
			}
		})
	}
}
