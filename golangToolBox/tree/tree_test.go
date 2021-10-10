package tree

import (
	"fmt"
	"testing"
)

func TestOrderRec(t *testing.T) {
	root := BuildBinaryTree()
	PreOrderRec(root)
	println()
	fmt.Println(PreOrderRecSlice(root))
	PreOrderUnRec(root)
	println()
	InOrderRec(root)
	println()
	InOrderUnRec(root)
	println()
	PostOrderRec(root)
	println()
	PostOrderUnRec(root)
	println()
	PostOrderUnRec2(root)
	println()
}
