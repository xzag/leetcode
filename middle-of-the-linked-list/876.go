package middle_of_the_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	single := head
	double := head

	for single != nil && double != nil && double.Next != nil {
		single = single.Next
		double = double.Next.Next
	}

	return single
}
