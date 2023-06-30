package middle_of_the_liked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	max := 1
	cur := head

	for cur != nil {
		cur = cur.Next
		max++
	}

	middle := max/2 + max%2
	cur = head

	for i := 1; i < middle; i++ {
		cur = cur.Next
	}

	return cur
}
