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

func PreOrderRecSlice(root *TreeNode) []int {
	var ret []int
	var order func(node *TreeNode)
	order = func(node *TreeNode) {
		if node == nil {
			return
		}
		ret = append(ret, node.Value)
		order(node.Left)
		order(node.Right)
	}
	order(root)
	return ret
}

// PreOrderUnRec 中左右
// 当前节点，先压右孩子，再压左孩子
// 弹出，打印
func PreOrderUnRec(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		head := stack[len(stack)-1]
		fmt.Printf("%d->", head.Value)
		stack = stack[:len(stack)-1]
		if head.Right != nil {
			stack = append(stack, head.Right)
		}
		if head.Left != nil {
			stack = append(stack, head.Left)
		}
	}
}

// InOrderUnRec 左中右
// 把左孩子全部入栈
// 当前节点为空，从栈中拿一个，打印，当前节点往右跑
// 当前节点不为空，压栈，当前节点继续往左跑
func InOrderUnRec(root *TreeNode) {
	if root == nil {
		return
	}
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		if root != nil {
			stack = append(stack, root)
			root = root.Left
		} else {
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			fmt.Printf("%d->", root.Value)
			root = root.Right
		}
	}
}

// PostOrderUnRec 左右中
// 通过两个栈实现
// 中右左 => 存栈 => 栈中顺序，左右中
// 第一个栈先进左孩子，再进右孩子，
func PostOrderUnRec(root *TreeNode) {
	if root == nil {
		return
	}
	stack1 := []*TreeNode{root}
	var stack2 []*TreeNode
	for len(stack1) > 0 {
		head := stack1[len(stack1)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = append(stack2, head)
		if head.Left != nil {
			stack1 = append(stack1, head.Left)
		}
		if head.Right != nil {
			stack1 = append(stack1, head.Right)
		}
	}
	for i := len(stack2) - 1; i >= 0; i-- {
		fmt.Printf("%d->", stack2[i].Value)
	}
}

// PostOrderUnRec2 通过一个栈实现 左右中
// 增加一个标记位，记录上次访问过的节点
func PostOrderUnRec2(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	last := root
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		if top.Left != nil && top.Left != last && top.Right != last {
			// 左孩子不空，且上次未访问过
			stack = append(stack, top.Left)
		} else if top.Right != nil && top.Right != last {
			// 右孩子不空，且上次未访问过
			stack = append(stack, top.Right)
		} else {
			fmt.Printf("%d->", top.Value)
			last = top
			stack = stack[:len(stack)-1]
		}
	}
}
