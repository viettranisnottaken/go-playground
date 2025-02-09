package two_pointer

import "testing"

type ThreeSumTestCase struct {
	numbers []int
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name   string
		inputs ThreeSumTestCase
		expect [][3]int
	}{
		{
			name: "Case 1",
			inputs: ThreeSumTestCase{
				numbers: []int{-1, 0, 1, 2, -1, -4},
			},
			expect: [][3]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "Case 2",
			inputs: ThreeSumTestCase{
				numbers: []int{0, 1, 1},
			},
			expect: [][3]int{},
		},
		{
			name: "Case 3",
			inputs: ThreeSumTestCase{
				numbers: []int{-2, 0, 0, 2, 2},
			},
			expect: [][3]int{
				{-2, 0, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ThreeSum(tt.inputs.numbers)

			isAllEqual := true

			for i, v := range got {
				if v != tt.expect[i] {
					isAllEqual = false

					break
				}
			}

			if len(got) != len(tt.expect) || !isAllEqual {
				t.Errorf("TwoIntegerSum2(%d) = %d; expected %d", tt.inputs.numbers, got, tt.expect)
			}
		})
	}
}
