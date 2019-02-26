package main

import "fmt"

func main() {
	fmt.Println(numJewelsInStones("aA", "aAAZZzz"))
}

func numJewelsInStones(J string, S string) int {
	// 申请一个map，统计出现过的字符(使用int32存储)
	jewels := make(map[rune]bool)
	for _, j := range J {
		jewels[j] = true
	}
	// 统计出现的次数
	count := 0
	for _, s := range S {
		if jewels[s] {
			count++
		}
	}
	return count
}
