package main

import "fmt"

func main() {
	nums := []int{1}
	fmt.Println(topKFrequent(nums, 1))
}

// top k appear frequent element
func topKFrequent(nums []int, k int) []int {
	// apply map to record element count
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[num] += 1
	}
	// record the frequentest count
	max := 0
	// apply map to record what elements are in each frequency
	temps := make(map[int][]int, len(nums))
	for k, v := range maps {
		if temps[v] == nil {
			temps[v] = make([]int, 0)
		}
		temps[v] = append(temps[v], k)
		if v > max {
			max = v
		}
	}
	// exit when it equal k
	count := 0
	// output
	ret := make([]int, 0)
	for i := max; i > 0; i-- {
		if temps[i] != nil {
			for _, num := range temps[i] {
				if count == k {
					return ret
				}
				ret = append(ret, num)
				count++
			}
		}
	}
	return ret
}
