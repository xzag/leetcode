package palindrome_linked_list

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{
			name: "simple no",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
				},
			},
			want: false,
		},
		{
			name: "double yes",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: true,
		},
		{
			name: "double",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val:  1,
							Next: nil,
						},
					},
				},
			},
			want: false,
		},
		{
			name: "single yes",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
					},
				},
			},
			want: true,
		},
		{
			name: "steps",
			head: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 1,
							},
						},
					},
				},
			},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, isPalindrome(tt.head))
		})
	}
}
