package problem0435

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(intervals [][]int) int) {
	// 测试用例。
	testCases := []struct {
		intervals [][]int
		want      int
	}{
		{
			intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}},
			want:      1,
		},
		{
			intervals: [][]int{{1, 2}, {1, 2}, {1, 2}},
			want:      2,
		},
		{
			intervals: [][]int{{1, 2}, {2, 3}},
			want:      0,
		},
		{
			intervals: [][]int{{1, 100}, {11, 22}, {1, 11}, {2, 12}},
			want:      2,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.intervals)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestEraseOverlapIntervals(t *testing.T) {
	testFramework(t, eraseOverlapIntervals)
}

func TestEraseOverlapIntervals_greedy(t *testing.T) {
	testFramework(t, eraseOverlapIntervals_greedy)
}
