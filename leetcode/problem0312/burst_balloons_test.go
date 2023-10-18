package problem0312

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{3, 1, 5, 8},
			want: 167,
		},
		{
			nums: []int{1, 5},
			want: 10,
		},
		{
			nums: []int{7, 9, 8, 0, 7, 1, 3, 5, 5, 2, 3},
			want: 1654,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaxCoins_dynamic_programming(t *testing.T) {
	testFramework(t, maxCoins_dynamic_programming)
}

func TestMaxCoins_memorized(t *testing.T) {
	testFramework(t, maxCoins_memorized)
}

func TestMaxCoins_brute(t *testing.T) {
	testFramework(t, maxCoins_brute)
}
