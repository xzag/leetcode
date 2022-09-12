package solutions

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_sample(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{
			name:  "sample",
			input: "123",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, sample(tt.input))
		})
	}
}
