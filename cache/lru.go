package cache

type LRU struct {
	Head *ListNode
	Tail *ListNode
	m    map[int]*ListNode
}

type ListNode struct {
	Prev *ListNode
	Next *ListNode
	Val  int
}

func NewListNode() (*ListNode, *ListNode) {
	head := new(ListNode)
	tail := new(ListNode)
	head.Next = tail
	tail.Prev = head
	return head, tail
}

func NewLRU() *LRU {
	head, tail := NewListNode()
	return &LRU{
		Head: head,
		Tail: tail,
		m:    make(map[int]*ListNode),
	}
}

func (l *ListNode) Insert(node *ListNode) {
	node.Prev = l
	node.Next = l.Next
	l.Next.Prev = node
	l.Next = node
}

func (l *ListNode) Remove() *ListNode {
	l.Prev.Next = l.Next
	l.Next.Prev = l.Prev
	l.Prev = nil
	l.Next = nil
	return l
}

func (l *LRU) Get(key int) int {
	node, ok := l.m[key]
	if !ok {
		return -1
	}
	node = node.Remove()
	l.Head.Insert(node)
	return node.Val
}

func (l *LRU) Set(key int, val int) {
	node, ok := l.m[key]
	if !ok {
		node = &ListNode{Val: val}
		l.m[key] = node
	} else {
		node.Remove()
	}
	l.Head.Insert(node)
}
