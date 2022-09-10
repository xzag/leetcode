package add_two_numbers

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		{
			name: "no carry",
			l1: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			l2: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
				},
			},
		},
		{
			name: "carry",
			l1: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 2,
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
		{
			name: "no carry different length",
			l1: &ListNode{
				Val: 6,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			l2: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 1,
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, addTwoNumbers(tt.l1, tt.l2))
		})
	}
}
