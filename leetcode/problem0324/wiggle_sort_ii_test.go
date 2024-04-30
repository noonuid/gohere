package problem0324

import (
	"reflect"
	"sort"
	"testing"
)

func test(t *testing.T, testFunc func(nums []int)) {
	// 测试用例。
	testCases := []struct {
		nums   []int
		sorted []int
	}{
		{
			nums:   []int{1, 5, 1, 1, 6, 4},
			sorted: []int{1, 1, 1, 4, 5, 6},
		},
		{
			nums:   []int{1, 3, 2, 2, 3, 1},
			sorted: []int{1, 1, 2, 2, 3, 3},
		},
		{
			nums:   []int{1, 1, 2, 1, 2, 2, 1},
			sorted: []int{1, 1, 1, 1, 2, 2, 2},
		},
		{
			nums:   []int{4, 5, 5, 6},
			sorted: []int{4, 5, 5, 6},
		},
	}

	isValid := func(nums []int) bool {
		for i := 0; i < len(nums)-1; i++ {
			if i%2 == 0 {
				if nums[i] >= nums[i+1] {
					return false
				}
			} else {
				if nums[i] <= nums[i+1] {
					return false
				}
			}
		}
		return true
	}

	for caseIndex, testCase := range testCases {
		testFunc(testCase.nums)
		numsCopy := append([]int{}, testCase.nums...)
		sort.Ints(numsCopy)
		if !isValid(testCase.nums) || !reflect.DeepEqual(numsCopy, testCase.sorted) {
			t.Errorf("\ncaseIndex: %d\ngot: %v",
				caseIndex, testCase.nums)
		}
	}
}

func TestWiggleSort(t *testing.T) {
	test(t, wiggleSort)
}
