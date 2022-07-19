package travel

import (
	"testing"

	"github.com/dinghenc/leetcode-go/tree"
	"github.com/stretchr/testify/assert"
)

func TestOrderDfs(t *testing.T) {
	s := "1#2#4#nil#nil#5#nil#6#nil#nil#3#7#8#nil#nil#nil#nil#"
	root := tree.Unmarshal(s)

	preorderTravel := PreOrder(root)
	inorderTravel := InOrder(root)
	postorderTravel := PostOrder(root)

	preorderTravelWithStack := PreOrderWithStack(root)
	inorderTravelWithStack := InOrderWithStack(root)
	postorderTravelWithStack := PostOrderWithStack(root)

	t.Log("recursive")
	t.Logf("preorder: %v\n", preorderTravel)
	t.Logf("inorder: %v\n", inorderTravel)
	t.Logf("postorder: %v\n", postorderTravel)

	t.Log("stack")
	t.Logf("preorder: %v\n", preorderTravelWithStack)
	t.Logf("inorder: %v\n", inorderTravelWithStack)
	t.Logf("postorder: %v\n", postorderTravelWithStack)

	assert.Equal(t, preorderTravel, preorderTravelWithStack)
	assert.Equal(t, inorderTravel, inorderTravelWithStack)
	assert.Equal(t, postorderTravel, postorderTravelWithStack)
}
