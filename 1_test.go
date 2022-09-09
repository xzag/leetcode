package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		target int
		want   []int
	}{
		{
			name:   "sorted",
			input:  []int{2, 4, 5},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "sorted 2",
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "not sorted 2",
			input:  []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "same",
			input:  []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "negative, sum positive",
			input:  []int{-3, -10, 7, 4},
			target: 4,
			want:   []int{0, 2},
		},
		{
			name:   "all negative",
			input:  []int{-3, -1, -7, -4},
			target: -5,
			want:   []int{1, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, twoSum(tt.input, tt.target))
			require.Equal(t, tt.want, twoSum_WithSort(tt.input, tt.target))
		})
	}
}
