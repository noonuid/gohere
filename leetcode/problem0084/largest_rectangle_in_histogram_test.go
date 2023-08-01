package problem0084

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例
	testCases := []struct {
		heights []int
		want    int
	}{
		{
			heights: []int{2, 1, 5, 6, 2, 3},
			want:    10,
		},
		{
			heights: []int{2, 4},
			want:    4,
		},
		{
			heights: []int{1},
			want:    1,
		},
		{
			heights: []int{0, 9},
			want:    9,
		},
		{
			heights: []int{4, 2},
			want:    4,
		},
		{
			heights: []int{1, 10, 10, 10, 2, 2, 3},
			want:    30,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.heights)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLargestRectangleArea_monotone_stack(t *testing.T) {
	testFramework(t, largestRectangleArea_monotone_stack)
}

func TestLargestRectangleArea_monotone_stack_sentinel(t *testing.T) {
	testFramework(t, largestRectangleArea_monotone_stack_sentinel)
}

func TestLargestRectangleArea_brute_height(t *testing.T) {
	testFramework(t, largestRectangleArea_brute_height)
}

func TestLargestRectangleArea_brute_left_right(t *testing.T) {
	testFramework(t, largestRectangleArea_brute_left_right)
}
