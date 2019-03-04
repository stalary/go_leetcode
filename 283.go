package main

import "fmt"

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}

// move 0 to end
func moveZeroes(nums []int) {
	// len
	l := len(nums)
	// record index
	pos := 0
	for _, n := range nums {
		// not 0 write array
		if n != 0 {
			nums[pos] = n
			pos++
		}
	}
	// write 0
	for i := pos; i < l; i++ {
		nums[i] = 0
	}
}
