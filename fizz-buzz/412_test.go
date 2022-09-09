package fizz_buzz

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_fizzBuzz(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want []string
	}{
		{
			name: "no",
			n:    2,
			want: []string{"1", "2"},
		},
		{
			name: "only fizz",
			n:    3,
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "fizz buzz",
			n:    5,
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "fizzBuzz",
			n:    15,
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, fizzBuzz(tt.n))
		})
	}
}
