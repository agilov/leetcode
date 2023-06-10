package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := ListNode{0, nil}
	add(l1, l2, &result, 0)
	return &result
}

func add(l1 *ListNode, l2 *ListNode, result *ListNode, carry int) {
	if l1 == nil && l2 == nil && carry == 0 {
		return
	}

	v1 := 0
	v2 := 0

	if l1 != nil {
		v1 = l1.Val
	}

	if l2 != nil {
		v2 = l2.Val
	}

	result.Val = v1 + v2 + carry
	if result.Val >= 10 {
		carry = 1
		result.Val -= 10
	} else {
		carry = 0
	}

	var next1 *ListNode
	var next2 *ListNode

	if l1 != nil {
		next1 = l1.Next
	}

	if l2 != nil {
		next2 = l2.Next
	}

	if next1 == nil && next2 == nil && carry == 0 {
		return
	}

	next := ListNode{0, nil}
	result.Next = &next

	add(next1, next2, &next, carry)
}
