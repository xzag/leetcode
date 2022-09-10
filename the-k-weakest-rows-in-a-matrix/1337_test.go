package the_k_weakest_rows_in_a_matrix

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_kWeakestRows(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		k    int
		want []int
	}{
		{
			name: "two",
			mat:  [][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}},
			k:    3,
			want: []int{2, 0, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, kWeakestRows(tt.mat, tt.k))
		})
	}
}
