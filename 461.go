package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

func hammingDistance(x int, y int) int {
	// 异或操作
	or := x ^ y
	// 获取总共有多少个1
	count := bits.OnesCount(uint(or))
	return count
}
