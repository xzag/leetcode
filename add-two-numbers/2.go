package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	var carry int
	sum := l1.Val + l2.Val
	if sum > 9 {
		carry = 1
		sum = sum - 10
	} else {
		carry = 0
	}

	next := addTwoNumbers(l1.Next, l2.Next)
	if carry > 0 {
		next = addTwoNumbers(next, &ListNode{Val: carry})
	}

	head := &ListNode{
		Val:  sum,
		Next: next,
	}

	return head
}
