package tree

import "testing"

func TestOrderRec(t *testing.T) {
	root := BuildBinaryTree()
	PreOrderRec(root)
	println()
	InOrderRec(root)
	println()
	PostOrderRec(root)
}
