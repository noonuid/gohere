package problem0027

import (
	"reflect"
	"sort"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int, val int) int) {
	testCases := []struct {
		nums []int
		val  int
		k    int
		want []int
	}{
		{
			nums: []int{3, 2, 2, 3},
			val:  3,
			k:    2,
			want: []int{2, 2},
		},
		{
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			k:    5,
			want: []int{0, 0, 1, 3, 4},
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums, testCase.val)
		if !reflect.DeepEqual(got, testCase.k) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.k)
		}

		sort.Ints(testCase.nums[:got])
		if !reflect.DeepEqual(testCase.nums[:got], testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.nums[:got], testCase.want)
		}
	}
}

func TestRemoveElement_opposite_pointer(t *testing.T) {
	testFramework(t, removeElement_opposite_pointer)
}

func TestRemoveElement_two_pointer(t *testing.T) {
	testFramework(t, removeElement_two_pointer)
}
