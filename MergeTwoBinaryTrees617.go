package main

import "fmt"

func main() {
	t1 := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	t2 := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(mergeTrees(t1, t2))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	// 对当前节点求和
	t1.Val += t2.Val
	// 递归左节点
	t1.Left = mergeTrees(t1.Left, t2.Left)
	// 递归右节点
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
