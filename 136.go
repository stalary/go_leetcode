package main

import "fmt"

func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums))
}

// count single number, other appear twice
func singleNumber(nums []int) int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	return xor
}
