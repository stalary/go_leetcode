package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAZZzz"))
}

// s contains how many j
func numJewelsInStones(J string, S string) int {
	// statistic char
	jewels := make(map[rune]bool)
	for _, j := range J {
		jewels[j] = true
	}
	// statistic appear count
	count := 0
	for _, s := range S {
		if jewels[s] {
			count++
		}
	}
	return count
}
