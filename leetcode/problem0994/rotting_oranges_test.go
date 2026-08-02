package problem0994

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(grid [][]int) int) {
	// 测试用例。
	testCases := []struct {
		grid [][]int
		want int
	}{
		{
			grid: [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			want: 4,
		},
		{
			grid: [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			want: -1,
		},
		{
			grid: [][]int{{0, 2}},
			want: 0,
		},
		{
			grid: [][]int{{0}},
			want: 0,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.grid)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestOrangesRotting(t *testing.T) {
	testFramework(t, orangesRotting)
}
