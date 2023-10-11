package reverse_linked_2

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	var (
		counter                  uint16 = 1
		left8                           = uint16(left)
		right8                          = uint16(right)
		curr                            = head
		out, start, start2, prev *ListNode
	)

	if curr.Next == nil {
		return curr
	}

	if left == 1 {
		start2 = curr
		goto L
	}

	out = curr
	for counter < left8-1 {
		curr = curr.Next
		counter++
	}
	start = curr

	curr = curr.Next
	counter++

	start2 = curr

L:
	prev = curr
	curr = curr.Next
	counter++

	for counter <= right8 {
		curr.Next, curr, prev = prev, curr.Next, curr
		counter++
	}

	if start != nil {
		start.Next = prev
	} else {
		out = prev
	}

	if curr != nil {
		start2.Next = curr
	} else {
		start2.Next = nil
	}
	return out

}
