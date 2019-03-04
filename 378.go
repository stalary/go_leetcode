package main

import "fmt"

func main() {
	matrix := [][]int{{1, 2}, {1, 3}}
	fmt.Println(kthSmallest(matrix, 2))
}

// find matrix kth smallest number
func kthSmallest(matrix [][]int, k int) int {
	// the number of rows
	n := len(matrix)
	// low
	lo := matrix[0][0]
	// high
	hi := matrix[n-1][n-1]
	var mid, count int

	for lo < hi {
		// mid
		mid = lo + (hi-lo)/2
		// count less or equal mid
		count = countLEQ(matrix, mid)

		// min++ when count less than k
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
