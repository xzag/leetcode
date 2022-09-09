package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_romanToInt(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "M",
			want:  1000,
		},
		{
			input: "XVI",
			want:  16,
		},
		{
			input: "IX",
			want:  9,
		},
		{
			input: "LVIII",
			want:  58,
		},
		{
			input: "MCMXCIV",
			want:  1994,
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			require.Equal(t, tt.want, romanToInt(tt.input))
		})
	}
}
