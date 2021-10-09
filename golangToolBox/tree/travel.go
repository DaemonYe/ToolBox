package tree

import "fmt"

func PreOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d->", root.Value)
	PreOrderRec(root.Left)
	PreOrderRec(root.Right)
}

func InOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderRec(root.Left)
	fmt.Printf("%d->", root.Value)
	InOrderRec(root.Right)
}

func PostOrderRec(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderRec(root.Left)
	PostOrderRec(root.Right)
	fmt.Printf("%d->", root.Value)
}
