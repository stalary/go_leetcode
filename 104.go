package main

import "math"

func main() {

}

// 计算二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 计算左节点深度
	left := maxDepth(root.Left)
	// 计算右节点深度
	right := maxDepth(root.Right)
	// 将左右两边较大值相加
	return int(math.Max(float64(left), float64(right))) + 1
}
