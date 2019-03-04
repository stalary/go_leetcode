package main

import "stalary/study"

func main() {

}

// exchange tree left and right node
func invertTree(root *study.TreeNode) *study.TreeNode {
	if root == nil {
		return nil
	}
	// exchange cur node
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	// exchange child node
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
