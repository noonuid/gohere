package problem0189

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int, k int)) {
	// 测试用例。
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			k:    3,
			want: []int{5, 6, 7, 1, 2, 3, 4},
		},
		{
			nums: []int{-1, -100, 3, 99},
			k:    2,
			want: []int{3, 99, -1, -100},
		},
		{
			nums: []int{-1},
			k:    2,
			want: []int{-1},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6},
			k:    4,
			want: []int{3, 4, 5, 6, 1, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		testFunc(testCase.nums, testCase.k)
		if !reflect.DeepEqual(testCase.nums, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.nums, testCase.want)
		}
	}
}

func TestRotate_reverse(t *testing.T) {
	testFramework(t, rotate_reverse)
}

func TestRotate_circular_replacement(t *testing.T) {
	testFramework(t, rotate_circular_replacement)
}

func TestRotate_copy(t *testing.T) {
	testFramework(t, rotate_copy)
}
