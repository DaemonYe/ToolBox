package tree

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

// BuildBinaryTree 初始化树
func BuildBinaryTree() *TreeNode {
	root := &TreeNode{1, nil, nil}
	l := &TreeNode{2, nil, nil}
	r := &TreeNode{3, nil, nil}
	root.Left = l
	root.Right = r
	l2 := &TreeNode{4, nil, nil}
	ll2 := &TreeNode{5, nil, nil}
	l.Left = l2
	l.Right = ll2
	r2 := &TreeNode{6, nil, nil}
	rr2 := &TreeNode{7, nil, nil}
	r.Left = r2
	r.Right = rr2
	return root
}
