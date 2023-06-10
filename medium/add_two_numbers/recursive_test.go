package add_two_numbers

import "testing"

// go test -v ./medium/add_two_numbers/recursive*.go

// go test -v -run TestAddTwoNumbers ./medium/add_two_numbers/recursive*.go
func TestAddTwoNumbers(t *testing.T) {
	l1 := ListNode{1, &ListNode{1, nil}}
	l2 := ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}

	result := AddTwoNumbers(&l1, &l2)

	if result.Val != 2 {
		t.Errorf("AddTwoNumbers(%v, %v) first expected `%v`, but got `%v`", l1, l2, 2, result.Val)
	}

	if result.Next.Val != 2 {
		t.Errorf("AddTwoNumbers(%v, %v) second expected `%v`, but got `%v`", l1, l2, 2, result.Next.Val)
	}
}
