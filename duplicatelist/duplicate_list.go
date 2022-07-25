package duplicatelist

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(arr []int) *ListNode {
	head := new(ListNode)
	node := head
	for _, v := range arr {
		node.Next = &ListNode{Val: v}
		node = node.Next
	}
	return head.Next
}

func (l *ListNode) Slice() []int {
	var v []int
	for node := l; node != nil; node = node.Next {
		v = append(v, node.Val)
	}
	return v
}

func deleteDuplicates(head *ListNode) *ListNode {
	newHead := &ListNode{Next: head}
	for curr := newHead; curr.Next != nil && curr.Next.Next != nil; {
		if curr.Next.Val == curr.Next.Next.Val {
			v := curr.Next.Val
			for curr.Next != nil && curr.Next.Val == v {
				curr.Next = curr.Next.Next
			}
		} else {
			curr = curr.Next
		}
	}
	return newHead.Next
}

func deleteDuplicates2(head *ListNode) *ListNode {
	for curr := head; curr != nil; {
		for curr.Next != nil && curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		}
		curr = curr.Next
	}
	return head
}
