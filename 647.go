package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
}

// 统计回文子串的数量
func countSubstrings(s string) int {
	total := 0
	for mid := 0; mid < len(s); mid++ {
		// 判断多个元素的
		left := mid
		right := mid + 1
		total = help(s, left, right, total)
		// 判断单个元素
		left = mid
		right = mid
		total = help(s, left, right, total)
	}
	return total
}

func help(s string, left int, right int, total int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
		total++
	}
	return total
}
