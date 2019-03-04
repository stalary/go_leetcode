package main

import (
	"math"
	"stalary/study"
)

func main() {

}

// calculate tree depth
func maxDepth(root *study.TreeNode) int {
	if root == nil {
		return 0
	}
	// left depth
	left := maxDepth(root.Left)
	// right depth
	right := maxDepth(root.Right)
	// sum max(left, right)
	return int(math.Max(float64(left), float64(right))) + 1
}
