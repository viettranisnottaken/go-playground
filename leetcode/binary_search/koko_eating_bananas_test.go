package binary_search

import "testing"

type KokoEatingBananasTestInput struct {
	piles []int
	hours int
}

func TestKokoEatingBananas(t *testing.T) {
	tests := []struct {
		name   string
		inputs KokoEatingBananasTestInput
		expect int
	}{
		{
			name: "Case 1: hours exists",
			inputs: KokoEatingBananasTestInput{
				piles: []int{1, 4, 3, 2},
				hours: 9,
			},
			expect: 2,
		},
		{
			name: "Case 2: hours does not exist",
			inputs: KokoEatingBananasTestInput{
				piles: []int{25, 10, 23, 4},
				hours: 4,
			},
			expect: 25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := KokoEatingBananas(tt.inputs.piles, tt.inputs.hours)

			if got != tt.expect {
				t.Errorf("KokoEatingBananas(%v, %d) = %d; expected %d", tt.inputs.piles, tt.inputs.hours, got, tt.expect)
			}
		})
	}
}
