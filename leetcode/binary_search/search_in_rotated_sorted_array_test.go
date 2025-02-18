package binary_search

import "testing"

type SearchInRotatedSortedArrayTestInput struct {
	nums   []int
	target int
}

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		inputs SearchInRotatedSortedArrayTestInput
		expect int
	}{
		{
			name: "Case 1: target in the middle",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{4, 5, 6, 0, 1, 2, 3},
				target: 0,
			},
			expect: 3,
		},
		{
			name: "Case 2",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{3, 4, 5, 6, 1, 2},
				target: 1,
			},
			expect: 4,
		},
		{
			name: "Case 3",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{3, 5, 6, 0, 1, 2},
				target: 4,
			},
			expect: -1,
		},
		{
			name: "Case 4",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{5, 1, 3},
				target: 5,
			},
			expect: 0,
		},
		{
			name: "Case 5",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{1, 3},
				target: 3,
			},
			expect: 1,
		},
		{
			name: "Case 6",
			inputs: SearchInRotatedSortedArrayTestInput{
				nums:   []int{3, 5, 1},
				target: 3,
			},
			expect: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchInRotatedSortedArray(tt.inputs.nums, tt.inputs.target)

			if got != tt.expect {
				t.Errorf("SearchInRotatedSortedArray(%v, %d) = %d; expected %d", tt.inputs.nums, tt.inputs.target, got, tt.expect)
			}
		})
	}
}
