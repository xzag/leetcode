package palindrome_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var current, prev, next *ListNode

	current = head
	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}

func isPalindrome(head *ListNode) bool {
	single := head
	double := head
	var prev *ListNode
	for single != nil && double != nil && double.Next != nil {
		prev = single
		single = single.Next
		double = double.Next.Next
	}

	if double != nil {
		prev = single
	}

	reversed := reverseList(single)

	if prev != nil {
		prev.Next = nil
	}

	for head != nil && reversed != nil {
		if head.Val != reversed.Val {
			return false
		}

		head = head.Next
		reversed = reversed.Next
	}

	return head == reversed
}
