package problem0198

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			nums: []int{1, 3, 1},
			want: 3,
		},
		{
			nums: []int{2, 1, 1, 2},
			want: 4,
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

func TestRob(t *testing.T) {
	testFramework(t, rob)
}
