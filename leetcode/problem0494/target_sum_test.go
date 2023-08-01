package problem0494

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) int) {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{1, 1, 1, 1, 1},
			target: 3,
			want:   5,
		},
		{
			nums:   []int{1},
			target: 1,
			want:   1,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.nums, testCase.target)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindTargetSumWays_backtrack(t *testing.T) {
	testFramework(t, findTargetSumWays_backtrack)
}

func TestFindTargetSumWays_dynamic_programming(t *testing.T) {
	testFramework(t, findTargetSumWays_dynamic_programming)
}