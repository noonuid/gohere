package problem0034

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int, target int) []int) {
	// 测试用例。
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums, testCase.target)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestSearchRange(t *testing.T) {
	testFramework(t, searchRange)
}
