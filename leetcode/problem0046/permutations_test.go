package problem0046

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) [][]int) {
	// 测试用例
	testCases := []struct {
		nums []int
		want [][]int
	}{
		{
			nums: []int{1, 2, 3},
			want: [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			nums: []int{0, 1},
			want: [][]int{{0, 1}, {1, 0}},
		},
		{
			nums: []int{1},
			want: [][]int{{1}},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestPermute(t *testing.T) {
	testFramework(t, permute)
}
