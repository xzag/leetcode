package number_of_steps_to_reduce_a_number_to_zero

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_numberOfSteps(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{
			name: "even",
			n:    8,
			want: 4,
		},
		{
			name: "odd",
			n:    123,
			want: 12,
		},
		{
			name: "even more",
			n:    14,
			want: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, numberOfSteps(tt.n))
		})
	}
}
