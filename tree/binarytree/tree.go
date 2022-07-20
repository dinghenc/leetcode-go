package binarytree

import (
	"strconv"
	"strings"
)

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func New(val int) *Node {
	return &Node{Val: val}
}

func Marshal(root *Node) string {
	if root == nil {
		return "nil#"
	}
	var b strings.Builder
	b.WriteString(strconv.Itoa(root.Val) + "#")
	b.WriteString(Marshal(root.Left))
	b.WriteString(Marshal(root.Right))
	return b.String()
}

func Unmarshal(s string) *Node {
	if s == "" {
		return nil
	}
	nodeStrArr := strings.Split(s, "#")
	index := 0
	var dfs func() *Node
	dfs = func() *Node {
		if nodeStrArr[index] == "nil" {
			return nil
		}
		val, _ := strconv.Atoi(nodeStrArr[index])
		root := New(val)
		index++
		root.Left = dfs()
		index++
		root.Right = dfs()
		return root
	}
	return dfs()
}

func Equal(r1, r2 *Node) bool {
	if r1 == nil && r2 == nil {
		return true
	} else if r1 == nil || r2 == nil {
		return false
	} else if r1.Val != r2.Val {
		return false
	}
	return Equal(r1.Left, r2.Left) && Equal(r1.Right, r2.Right)
}
