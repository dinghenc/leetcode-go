package duplicatelist

import "testing"

func TestDeleteDuplicates(t *testing.T) {
	cases := [][]int{
		{1, 2, 3, 3, 4, 4, 5},
		{1, 1, 1, 2, 3},
		{1, 1, 1, 2, 3, 3},
	}
	for _, c := range cases {
		head := New(c)
		t.Log(head.Slice())
		head = deleteDuplicates(head)
		t.Log(head.Slice())
	}
}
