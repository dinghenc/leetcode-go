package cache

type LRUCache struct {
	capacity   int
	head, tail *ListNode
	m          map[int]*ListNode
}

func NewLRUCache(capacity int) LRUCache {
	head, tail := NewNode()
	return LRUCache{
		capacity: capacity,
		head:     head,
		tail:     tail,
		m:        make(map[int]*ListNode),
	}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.m[key]
	if !ok {
		return -1
	}
	node = node.Remove()
	c.head.Insert(node)
	return node.Val
}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.m[key]
	if ok {
		node = node.Remove()
		node.Val = value
	} else {
		if c.capacity == len(c.m) {
			del := c.tail.prev.Remove()
			delete(c.m, del.Key)
		}
		node = &ListNode{Key: key, Val: value}
		c.m[key] = node
	}
	c.head.Insert(node)
}

type ListNode struct {
	Key, Val   int
	prev, next *ListNode
}

func NewNode() (*ListNode, *ListNode) {
	head := new(ListNode)
	tail := new(ListNode)
	head.next = tail
	tail.prev = head
	return head, tail
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
	l.next = nil
	l.prev = nil
	return l
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
