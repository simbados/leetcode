package main

import (
	"math/rand"
)

type RandomizedSet struct {
	mapping map[int]int
	arr     []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		mapping: make(map[int]int),
		arr:     []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.mapping[val]; exists {
		return false
	}
	this.arr = append(this.arr, val)
	this.mapping[val] = len(this.arr) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, exists := this.mapping[val]
	if !exists {
		return false
	} else {
		last := this.arr[len(this.arr)-1]
		this.arr[index] = last
		this.mapping[last] = this.mapping[val]
		delete(this.mapping, val)
		this.arr = this.arr[0 : len(this.arr)-1]
		return true
	}
}

func (this *RandomizedSet) GetRandom() int {
	return this.arr[rand.Intn(len(this.arr))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
