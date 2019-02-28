package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// 将0移动到最后
func moveZeroes(nums []int) {
	// 记录数组长度
	l := len(nums)
	// 记录移动到的下标
	pos := 0
	for _, n := range nums {
		// 不等于0时写入数组
		if n != 0 {
			nums[pos] = n
			pos++
		}
	}
	// 写入等于0的元素
	for i := pos; i < l; i++ {
		nums[i] = 0
	}
}
