package travel

import (
	"github.com/dinghenc/leetcode-go/tree/binarytree"
)

func PreOrderWithStack(root *binarytree.Node) []int {
	var t []int
	var stack []*binarytree.Node
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			t = append(t, node.Val)
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = node.Right
	}
	return t
}

func InOrderWithStack(root *binarytree.Node) []int {
	var t []int
	var stack []*binarytree.Node
	node := root
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		t = append(t, node.Val)
		node = node.Right
	}
	return t
}

func PostOrderWithStack(root *binarytree.Node) []int {
	var t []int
	var stack []*binarytree.Node
	node := root
	var prev *binarytree.Node
	for node != nil || len(stack) != 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 右节点为空, 或者前一个访问的节点为右节点才能够访问当前节点
		if node.Right == nil || node.Right == prev {
			t = append(t, node.Val)
			prev = node
			// 根节点已经访问完毕, 需要置空使得左右子节点不再进入栈中
			node = nil
		} else {
			// 右节点还未访问, 先把当前节点压入栈中, 并访问右节点
			stack = append(stack, node)
			node = node.Right
		}
	}
	return t
}
