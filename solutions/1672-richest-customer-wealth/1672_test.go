package richest_customer_wealth

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		name     string
		accounts [][]int
		want     int
	}{
		{
			name:     "two",
			accounts: [][]int{{1, 2, 3}, {3, 2, 1}},
			want:     6,
		},
		{
			name:     "three",
			accounts: [][]int{{1, 5}, {7, 3}, {3, 5}},
			want:     10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, maximumWealth(tt.accounts))
		})
	}
}
