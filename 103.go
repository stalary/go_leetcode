package main

import (
	"stalary/study"
)

func main() {

}

// use 'z' print tree
func zigzagLevelOrder(root *study.TreeNode) [][]int {
	var result [][]int

	// determine direction
	flag := true
	// help save TreeNode
	stack := []*study.TreeNode{root}

	for len(stack) > 0 {
		var tmp []*study.TreeNode
		// cur level elements
		var levelVal []int
		for i := len(stack) - 1; i >= 0; i-- {
			node := stack[i]
			if node == nil {
				continue
			}
			if flag {
				tmp = append(tmp, node.Left, node.Right)
			} else {
				tmp = append(tmp, node.Right, node.Left)
			}
			levelVal = append(levelVal, node.Val)
		}
		flag = !flag
		stack = tmp
		if len(levelVal) > 0 {
			result = append(result, levelVal)
		}
	}
	return result
}
