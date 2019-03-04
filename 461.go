package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(hammingDistance(1, 4))
}

// calculate hamming distance
func hammingDistance(x int, y int) int {
	or := x ^ y
	// count one
	count := bits.OnesCount(uint(or))
	return count
}
