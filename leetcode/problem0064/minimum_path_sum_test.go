package problem0064

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([][]int) int) {
	// 测试用例。
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
			want: 7,
		},
		{
			grid: [][]int{{1, 2, 3}, {4, 5, 6}},
			want: 12,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.grid)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMinPathSum_dp_scrolling_array(t *testing.T) {
	testFramework(t, minPathSum_dp_scrolling_array)
}

func TestMinPathSum_dynamic_programming(t *testing.T) {
	testFramework(t, minPathSum_dynamic_programming)
}

func TestMinPathSum_backtrack(t *testing.T) {
	testFramework(t, minPathSum_backtrack)
}
