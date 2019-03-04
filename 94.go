package main

import (
	"container/list"
	"fmt"
	"stalary/study"
)

func main() {
	right := &study.TreeNode{
		Left: &study.TreeNode{
			Left:  nil,
			Right: nil,
			Val:   3,
		},
		Right: nil,
		Val:   2,
	}
	root := &study.TreeNode{
		Left:  nil,
		Right: right,
		Val:   1,
	}
	fmt.Println(inorderTraversal(root))
}

// inorder traversal tree
func inorderTraversal(root *study.TreeNode) []int {
	ret := make([]int, 0)
	if root == nil {
		return ret
	}
	stack := list.New()
	for root != nil || stack.Len() != 0 {
		for root != nil {
			fmt.Println(root.Val)
			stack.PushBack(root)
			root = root.Left
		}
		root = stack.Back().Value.(*study.TreeNode)
		ret = append(ret, root.Val)
		stack.Remove(stack.Back())
		root = root.Right
	}
	return ret
}
