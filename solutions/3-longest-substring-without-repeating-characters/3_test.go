package longest_substring_without_repeating_characters

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "startsWith",
			input: "abcabcbb",
			want:  3,
		},
		{
			name:  "same",
			input: "bbbbb",
			want:  1,
		},
		{
			name:  "middle",
			input: "pwwkew",
			want:  3,
		},
		{
			name:  "end",
			input: "aaaabcd",
			want:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, lengthOfLongestSubstring(tt.input))
		})
	}
}
