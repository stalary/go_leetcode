package main

import (
	"fmt"
	"stalary/study"
)

func main() {
	t1 := &study.TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	t2 := &study.TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	fmt.Println(mergeTrees(t1, t2))
}

// 合并两颗二叉树
func mergeTrees(t1 *study.TreeNode, t2 *study.TreeNode) *study.TreeNode {
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
