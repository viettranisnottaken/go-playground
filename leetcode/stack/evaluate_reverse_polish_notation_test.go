package stack

import "testing"

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{
			name:   "Basic operations",
			tokens: []string{"1", "2", "+", "3", "*", "4", "-"},
			want:   5, // ((1 + 2) * 3) - 4 = 5
		},
		{
			name:   "Handling negative numbers",
			tokens: []string{"-2", "3", "*", "4", "+"},
			want:   -2*3 + 4, // (-2 * 3) + 4 = -6 + 4 = -2
		},
		{
			name:   "Integer division truncation",
			tokens: []string{"10", "3", "/"},
			want:   10 / 3, // 10 / 3 = 3 (truncated towards zero)
		},
		{
			name:   "Handling large expressions",
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   4 + (13 / 5), // 4 + (13 / 5) = 4 + 2 = 6
		},
		{
			name:   "Handling single number",
			tokens: []string{"42"},
			want:   42, // Single number should be returned as is
		},
		{
			name:   "Complex expression",
			tokens: []string{"3", "4", "+", "2", "*", "7", "/"},
			want:   ((3 + 4) * 2) / 7, // (7 * 2) / 7 = 14 / 7 = 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := EvaluateReversePolishNotation(tt.tokens)
			if got != tt.want {
				t.Errorf("evalRPN(%v) = %d; want %d", tt.tokens, got, tt.want)
			}
		})
	}
}
