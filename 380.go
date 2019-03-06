package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	values []int
	index  map[int]int // record element index
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{[]int{}, make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	// return false if exist
	if _, exist := this.index[val]; exist {
		return false
	}
	this.values = append(this.values, val)
	// set index
	this.index[val] = len(this.values) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if i, exist := this.index[val]; !exist {
		return false
	} else {
		if len(this.values) < 2 {
			// return empty
			this.values = this.values[0:0]
		} else {
			// swap the last element
			swap := this.values[len(this.values)-1]
			this.index[swap] = i
			this.values[i] = swap
			this.values = this.values[0 : len(this.values)-1]
		}
		delete(this.index, val)
		return true
	}
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	return this.values[rand.Int()%len(this.index)]
}
