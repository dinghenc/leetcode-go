package cache

type LRU struct {
	head *ListNode
	tail *ListNode
	m    map[int]*ListNode
}

type ListNode struct {
	prev *ListNode
	next *ListNode
	Val  int
}

func NewListNode() (*ListNode, *ListNode) {
	head := new(ListNode)
	tail := new(ListNode)
	head.next = tail
	tail.prev = head
	return head, tail
}

func NewLRU() *LRU {
	head, tail := NewListNode()
	return &LRU{
		head: head,
		tail: tail,
		m:    make(map[int]*ListNode),
	}
}

func (l *ListNode) Insert(node *ListNode) {
	node.prev = l
	node.next = l.next
	l.next.prev = node
	l.next = node
}

func (l *ListNode) Remove() *ListNode {
	l.prev.next = l.next
	l.next.prev = l.prev
	l.prev = nil
	l.next = nil
	return l
}

func (l *LRU) Get(key int) int {
	node, ok := l.m[key]
	if !ok {
		return -1
	}
	node = node.Remove()
	l.head.Insert(node)
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
	l.head.Insert(node)
}
