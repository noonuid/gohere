package problem0121

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
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			prices: []int{2, 1, 2, 1, 0, 1, 2},
			want:   2,
		},
		{
			prices: []int{1, 2, 11, 4, 7},
			want:   10,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.prices)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMaxProfit(t *testing.T) {
	testFramework(t, maxProfit)
}

func TestMaxProfit_brute_force(t *testing.T) {
	testFramework(t, maxProfit_brute_force)
}
