package problem0077

import (
	"reflect"
	"sort"
	"testing"
)

func test(t *testing.T, fn func(n int, k int) [][]int) {
	testCases := []struct {
		n    int
		k    int
		want [][]int
	}{
		{
			n: 4,
			k: 2,
			want: [][]int{
				{2, 4},
				{3, 4},
				{2, 3},
				{1, 2},
				{1, 3},
				{1, 4},
			},
		},
		{
			n: 1,
			k: 1,
			want: [][]int{
				{1},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.n, testCase.k)

		sortAnswer := func(answer [][]int) {
			sort.Slice(answer, func(i, j int) bool {
				nums1, nums2 := answer[i], answer[j]
				minLen := len(nums1)
				if len(nums2) < len(nums1) {
					minLen = len(nums2)
				}
				for k := 0; k < minLen; k++ {
					if nums1[k] != nums2[k] {
						return nums1[k] < nums2[k]
					}
				}
				return len(nums1) < len(nums2)
			})
		}
		sortAnswer(got)
		sortAnswer(testCase.want)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCombine(t *testing.T) {
	test(t, combine)
}
