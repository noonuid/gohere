package problem0373

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(nums1 []int, nums2 []int, k int) [][]int) {
	testCases := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{
			nums1: []int{1, 7, 11},
			nums2: []int{2, 4, 6},
			k:     3,
			want:  [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			nums1: []int{1, 1, 2},
			nums2: []int{1, 2, 3},
			k:     2,
			want:  [][]int{{1, 1}, {1, 1}},
		},
		{
			nums1: []int{1, 2, 4, 5, 6},
			nums2: []int{3, 5, 7, 9},
			k:     3,
			want:  [][]int{{1, 3}, {2, 3}, {1, 5}},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.nums1, testCase.nums2, testCase.k)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestKSmallestPairs_heap(t *testing.T) {
	test(t, kSmallestPairs_heap)
}

func TestKSmallestPairs_enum(t *testing.T) {
	test(t, kSmallestPairs_enum)
}
