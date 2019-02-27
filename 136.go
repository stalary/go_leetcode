package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

// 一个数组中只有一个数出现一次，其他数都出现两次，找出出现一次的数
func singleNumber(nums []int) int {
	// 使用异或，最后的结果即为只出现一次的
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
