package two_pointer

import "testing"

type TrappingRainWaterTestCase struct {
	height []int
}

func TestTrappingRainWater(t *testing.T) {
	tests := []struct {
		name   string
		inputs TrappingRainWaterTestCase
		expect int
	}{
		{
			name: "Case 1",
			inputs: TrappingRainWaterTestCase{
				height: []int{0, 2, 0, 3, 1, 0, 1, 3, 2, 1},
			},
			expect: 9,
		},
		{
			name: "Case 2",
			inputs: TrappingRainWaterTestCase{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			expect: 6,
		},
		{
			name: "Case 3",
			inputs: TrappingRainWaterTestCase{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			expect: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TrappingRainWater(tt.inputs.height)

			if got != tt.expect {
				t.Errorf("TrappingRainWater(%d) = %d; expected %d", tt.inputs.height, got, tt.expect)
			}
		})
	}
}
