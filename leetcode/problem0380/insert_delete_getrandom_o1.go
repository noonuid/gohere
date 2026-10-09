package problem0380

import "math/rand"

type RandomizedSet struct {
	indices map[int]int
	nums    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		indices: make(map[int]int),
		nums:    []int{},
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.indices[val]; ok {
		return false
	}
	this.indices[val] = len(this.nums)
	this.nums = append(this.nums, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.indices[val]
	if !ok {
		return false
	}
	this.nums[index] = this.nums[len(this.nums)-1]
	this.indices[this.nums[index]] = index
	this.nums = this.nums[:len(this.nums)-1]
	delete(this.indices, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
