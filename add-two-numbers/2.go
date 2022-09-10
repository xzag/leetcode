package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var first, second, carry int

	fakeHead := &ListNode{}
	current := fakeHead
	for l1 != nil || l2 != nil || carry > 0 {
		first = 0
		second = 0

		if l1 != nil {
			first = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			second = l2.Val
			l2 = l2.Next
		}

		sum := first + second + carry
		carry = sum / 10
		sum = sum % 10

		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	return fakeHead.Next
}
