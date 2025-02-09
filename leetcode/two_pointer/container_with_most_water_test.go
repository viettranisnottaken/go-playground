package two_pointer

import "testing"

type ContainerWithMostWaterTestCase struct {
	heights []int
}

func TestContainerWithMostWater(t *testing.T) {
	tests := []struct {
		name   string
		inputs ContainerWithMostWaterTestCase
		expect int
	}{
		{
			name: "Case 1",
			inputs: ContainerWithMostWaterTestCase{
				heights: []int{1, 7, 2, 5, 4, 7, 3, 6},
			},
			expect: 36,
		},
		{
			name: "Case 2",
			inputs: ContainerWithMostWaterTestCase{
				heights: []int{2, 2, 2},
			},
			expect: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainerWithMostWater(tt.inputs.heights)

			if got != tt.expect {
				t.Errorf("TwoIntegerSum2(%d) = %d; expected %d", tt.inputs.heights, got, tt.expect)
			}
		})
	}
}
