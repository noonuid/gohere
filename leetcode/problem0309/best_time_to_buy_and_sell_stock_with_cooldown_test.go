package problem0309

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例
	testCases := []struct {
		prices []int
		want   int
	}{
		{
			prices: []int{1, 2, 3, 0, 2},
			want:   3,
		},
		{
			prices: []int{1},
			want:   0,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.prices)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	testFramework(t, maxProfit)
}
