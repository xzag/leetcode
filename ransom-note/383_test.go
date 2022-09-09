package ransom_note

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name     string
		note     string
		magazine string
		want     bool
	}{
		{
			name:     "single",
			note:     "a",
			magazine: "b",
			want:     false,
		},
		{
			name:     "easy",
			note:     "ab",
			magazine: "abc",
			want:     true,
		},
		{
			name:     "easy fail",
			note:     "ab",
			magazine: "ac",
			want:     false,
		},
		{
			name:     "duplicates",
			note:     "aa",
			magazine: "aacc",
			want:     true,
		},
		{
			name:     "duplicates fail",
			note:     "aa",
			magazine: "ac",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, canConstruct(tt.note, tt.magazine))
		})
	}
}
