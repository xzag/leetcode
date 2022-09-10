package best_time_to_buy_and_sell_stock

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "two",
			prices: []int{2, 4, 1},
			want:   2,
		},
		{
			name:   "complex",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "nothing",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maxProfit(tt.prices))
		})
	}
}
