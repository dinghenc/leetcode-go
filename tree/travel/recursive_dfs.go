package travel

import "github.com/dinghenc/leetcode-go/tree"

func PreOrder(root *tree.Node) []int {
	var t []int
	var dfs func(root *tree.Node)
	dfs = func(root *tree.Node) {
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

func InOrder(root *tree.Node) []int {
	var t []int
	var dfs func(root *tree.Node)
	dfs = func(root *tree.Node) {
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

func PostOrder(root *tree.Node) []int {
	var t []int
	var dfs func(root *tree.Node)
	dfs = func(root *tree.Node) {
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
