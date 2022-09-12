package middle_of_the_linked_list

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_middleNode(t *testing.T) {
	tests := []struct {
		name string
		init func() (*ListNode, *ListNode)
	}{
		{
			name: "1",
			init: func() (*ListNode, *ListNode) {
				middle := &ListNode{
					Val: 2,
				}

				return &ListNode{
					Val:  1,
					Next: middle,
				}, middle
			},
		},
		{

			name: "2",
			init: func() (*ListNode, *ListNode) {
				middle := &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				}

				return &ListNode{
					Val: 1,
					Next: &ListNode{
						Val:  2,
						Next: middle,
					},
				}, middle
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head, middle := tt.init()
			require.Equal(t, middle, middleNode(head))
		})
	}
}
