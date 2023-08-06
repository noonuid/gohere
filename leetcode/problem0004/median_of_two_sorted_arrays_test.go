package problem0004

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, []int) float64) {
	// 测试用例。
	testCases := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{
			nums1: []int{1, 3},
			nums2: []int{2},
			want:  2.0,
		},
		{
			nums1: []int{1, 2},
			nums2: []int{3, 4},
			want:  2.5,
		},
		{
			nums1: []int{},
			nums2: []int{2, 3},
			want:  2.5,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums1, testCase.nums2)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindMedianSortedArrays_half_traversal(t *testing.T) {
	testFramework(t, findMedianSortedArrays_half_traversal)
}

func TestFindMedianSortedArrays_merging(t *testing.T) {
	testFramework(t, findMedianSortedArrays_merging)
}
