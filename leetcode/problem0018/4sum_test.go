package problem0018

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) [][]int) {
	// 测试用例。
	testCases := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{
			nums:   []int{1, 0, -1, 0, -2, 2},
			target: 0,
			want:   [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			nums:   []int{2, 2, 2, 2, 2},
			target: 8,
			want:   [][]int{{2, 2, 2, 2}},
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

func TestFourSum_double_pointer(t *testing.T) {
	testFramework(t, fourSum_double_pointer)
}

func TestFourSum_brute(t *testing.T) {
	testFramework(t, fourSum_brute)
}
