package main

import "fmt"

func main() {
	fmt.Println(countSubstrings("aaa"))
}

// count palindrome substrings
func countSubstrings(s string) int {
	total := 0
	for mid := 0; mid < len(s); mid++ {
		// count multiple elements
		left := mid
		right := mid + 1
		total = help(s, left, right, total)
		// count one element
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
