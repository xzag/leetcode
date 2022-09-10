package palindrome_linked_list

/**
Given the head of a singly linked list, return true if it is a palindrome or false otherwise.

Example 1:

Input: head = [1,2,2,1]
Output: true

Example 2:

Input: head = [1,2]
Output: false

Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
*/

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
