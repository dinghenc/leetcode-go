package tree

import "testing"

func TestMarshal(t *testing.T) {
	root := New(1)
	root.Left = New(2)
	root.Right = New(3)
	root.Left.Left = New(4)
	root.Left.Right = New(5)
	root.Left.Right.Right = New(6)
	root.Right.Left = New(7)
	root.Right.Left.Left = New(8)

	s := Marshal(root)
	t.Log(s)
	copyRoot := Unmarshal(s)
	t.Log(Equal(root, copyRoot))
}
