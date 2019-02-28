package main

func main() {

}

// 交换二叉树左右节点
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	// 先对当前左右节点进行交换
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	// 递归交换子节点
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
