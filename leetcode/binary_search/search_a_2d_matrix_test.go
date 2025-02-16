package binary_search

import "testing"

type Search2DMatrixTestInput struct {
	matrix [][]int
	target int
}

func TestSearchA2DMatrix(t *testing.T) {
	tests := []struct {
		name   string
		inputs Search2DMatrixTestInput
		expect bool
	}{
		{
			name: "Case 1: target exists",
			inputs: Search2DMatrixTestInput{
				matrix: [][]int{
					{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40},
				},
				target: 10,
			},
			expect: true,
		},
		{
			name: "Case 2: target exists",
			inputs: Search2DMatrixTestInput{
				matrix: [][]int{
					{1, 2, 4, 8}, {10, 11, 12, 13}, {14, 20, 30, 40},
				},
				target: 14,
			},
			expect: true,
		},
		{
			name: "Case 3: target does not exist",
			inputs: Search2DMatrixTestInput{
				matrix: [][]int{
					{1, 2, 4, 8}, {10, 11, 1, 13}, {14, 20, 3, 40},
				},
				target: 15,
			},
			expect: false,
		},
		{
			name: "Case 4: out of bound",
			inputs: Search2DMatrixTestInput{
				matrix: [][]int{
					{1, 2, 4, 8}, {10, 11, 1, 13}, {14, 20, 3, 40},
				},
				target: 99,
			},
			expect: false,
		},
		{
			name: "Case 5: empty matrix",
			inputs: Search2DMatrixTestInput{
				matrix: [][]int{},
				target: 99,
			},
			expect: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchA2DMatrix(tt.inputs.matrix, tt.inputs.target)

			if got != tt.expect {
				t.Errorf("TwoIntegerSum2(%v, %d) = %t; expected %t", tt.inputs.matrix, tt.inputs.target, got, tt.expect)
			}
		})
	}
}
