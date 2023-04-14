package problem0300

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
			nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			want: 4,
		},
		{
			nums: []int{0, 1, 0, 3, 2, 3},
			want: 4,
		},
		{
			nums: []int{7, 7, 7, 7, 7, 7, 7},
			want: 1,
		},
		{
			nums: []int{4, 10, 4, 3, 8, 9},
			want: 3,
		},
		{
			nums: []int{1, 3, 6, 7, 9, 4, 10, 5, 6},
			want: 6,
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

func TestTwoSum(t *testing.T) {
	testFramework(t, lengthOfLIS)
}
