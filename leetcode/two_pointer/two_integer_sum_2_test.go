package two_pointer

import "testing"

type TwoSum2TestCase struct {
	numbers []int
	target  int
}

func TestTwoIntegerSum2(t *testing.T) {
	tests := []struct {
		name   string
		inputs TwoSum2TestCase
		expect [2]int
	}{
		{
			name: "Case 1",
			inputs: TwoSum2TestCase{
				numbers: []int{1, 2, 3, 4},
				target:  3,
			},
			expect: [2]int{1, 2},
		},
		{
			name: "Case 2: With negative",
			inputs: TwoSum2TestCase{
				numbers: []int{-1, 0},
				target:  -1,
			},
			expect: [2]int{1, 2},
		},
		{
			name: "Case 3: With negative",
			inputs: TwoSum2TestCase{
				numbers: []int{-5, -3, 0, 2, 4, 6, 8},
				target:  5,
			},
			expect: [2]int{2, 7},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoIntegerSum2(tt.inputs.numbers, tt.inputs.target)
			if got != tt.expect {
				t.Errorf("TwoIntegerSum2(%d, %d) = %d; expected %d", tt.inputs.numbers, tt.inputs.target, got, tt.expect)
			}
		})
	}
}
