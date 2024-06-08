package problem0088

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums1 []int, m int, nums2 []int, n int)) {
	testCases := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
		want  []int
	}{
		{
			nums1: []int{1, 2, 3, 0, 0, 0},
			m:     3,
			nums2: []int{2, 5, 6},
			n:     3,
			want:  []int{1, 2, 2, 3, 5, 6},
		},
		{
			nums1: []int{1},
			m:     1,
			nums2: []int{},
			n:     0,
			want:  []int{1},
		},
		{
			nums1: []int{0},
			m:     0,
			nums2: []int{1},
			n:     1,
			want:  []int{1},
		},
	}

	for caseIndex, testCase := range testCases {
		testFunc(testCase.nums1, testCase.m, testCase.nums2, testCase.n)
		if !reflect.DeepEqual(testCase.nums1, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.nums1, testCase.want)
		}
	}
}

func TestMerge(t *testing.T) {
	testFramework(t, merge)
}

func TestMerge_double_pointer(t *testing.T) {
	testFramework(t, merge_double_pointer)
}
