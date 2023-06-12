package middle_of_the_liked_list

import "testing"

func TestMiddleLinkedListNaive(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := middleNode(list)

	if result.Val != 3 {
		t.Error("Shoul be 3")
	}

	list = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = middleNode(list)

	if result.Val != 3 {
		t.Errorf("Should be 3 got %v", result.Val)
	}
}

func TestMiddleLinkedListFastAndSlow(t *testing.T) {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}
	result := middleNodeFastAndSlow(list)

	if result.Val != 3 {
		t.Error("Shoul be 3")
	}

	list = &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	result = middleNode(list)

	if result.Val != 3 {
		t.Errorf("Should be 3 got %v", result.Val)
	}
}
