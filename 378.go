package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2}, {1, 3}}
	fmt.Println(kthSmallest(matrix, 2))
}

// 查找矩阵中第k小的数字
func kthSmallest(matrix [][]int, k int) int {
	// 行数
	n := len(matrix)
	// 最小值
	lo := matrix[0][0]
	// 最大值
	hi := matrix[n-1][n-1]
	var mid, count int

	for lo < hi {
		// 求出中位数
		mid = lo + (hi-lo)/2
		// 计算小于中位数的数量
		count = countLEQ(matrix, mid)

		// 当数量低于k时，最小值增大
		if count < k {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return lo
}

func countLEQ(matrix [][]int, x int) int {
	n := len(matrix)
	count := 0
	var j int

	for _, row := range matrix {
		for j = 0; j < n && row[j] <= x; j++ {
			count++
		}
	}

	return count
}
