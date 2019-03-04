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

// merge two trees
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
	// sum cur node
	t1.Val += t2.Val
	// count left nodes
	t1.Left = mergeTrees(t1.Left, t2.Left)
	// count right nodes
	t1.Right = mergeTrees(t1.Right, t2.Right)
	return t1
}
