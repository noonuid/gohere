package problem0581

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{2, 6, 4, 8, 10, 9, 15},
			want: 5,
		},
		{
			nums: []int{1, 2, 3, 4},
			want: 0,
		},
		{
			nums: []int{1},
			want: 0,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindUnsortedSubarray_traversal(t *testing.T) {
	testFramework(t, findUnsortedSubarray_traversal)
}

func TestFindUnsortedSubarray_sort(t *testing.T) {
	testFramework(t, findUnsortedSubarray_sort)
}
