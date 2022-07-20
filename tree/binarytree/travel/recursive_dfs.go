package travel

import (
	"github.com/dinghenc/leetcode-go/tree/binarytree"
)

func PreOrder(root *binarytree.Node) []int {
	var t []int
	var dfs func(root *binarytree.Node)
	dfs = func(root *binarytree.Node) {
		if root == nil {
			return
		}
		t = append(t, root.Val)
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return t
}

func InOrder(root *binarytree.Node) []int {
	var t []int
	var dfs func(root *binarytree.Node)
	dfs = func(root *binarytree.Node) {
		if root == nil {
			return
		}
		dfs(root.Left)
		t = append(t, root.Val)
		dfs(root.Right)
	}
	dfs(root)
	return t
}

func PostOrder(root *binarytree.Node) []int {
	var t []int
	var dfs func(root *binarytree.Node)
	dfs = func(root *binarytree.Node) {
		if root == nil {
			return
		}
		dfs(root.Left)
		dfs(root.Right)
		t = append(t, root.Val)
	}
	dfs(root)
	return t
}
