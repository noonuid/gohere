package problem0152

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
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			nums: []int{-2, 0, -1},
			want: 0,
		},
		{
			nums: []int{-2},
			want: -2,
		},
		{
			nums: []int{-3, -1, -1},
			want: 3,
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

func TestMaxProduct_dynamic_programming(t *testing.T) {
	testFramework(t, maxProduct_dynamic_programming)
}

func TestMaxProduct(t *testing.T) {
	testFramework(t, maxProduct_brute)
}
