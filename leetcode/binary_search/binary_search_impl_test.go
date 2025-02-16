package binary_search

import "testing"

type BinarySearchImplTestInput struct {
	nums   []int
	target int
}

func TestBinarySearchImpl(t *testing.T) {
	tests := []struct {
		name   string
		inputs BinarySearchImplTestInput
		expect int
	}{
		{
			name: "Case 1: target exists",
			inputs: BinarySearchImplTestInput{
				nums:   []int{-1, 0, 2, 4, 6, 8},
				target: 4,
			},
			expect: 3,
		},
		{
			name: "Case 2: target does not exist",
			inputs: BinarySearchImplTestInput{
				nums:   []int{-1, 0, 2, 4, 6, 8},
				target: 3,
			},
			expect: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearchImpl(tt.inputs.nums, tt.inputs.target)

			if got != tt.expect {
				t.Errorf("TwoIntegerSum2(%v, %d) = %d; expected %d", tt.inputs.nums, tt.inputs.target, got, tt.expect)
			}
		})
	}
}
