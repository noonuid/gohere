package problem0041

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
			nums: []int{1, 2, 0},
			want: 3,
		},
		{
			nums: []int{3, 4, -1, 1},
			want: 2,
		},
		{
			nums: []int{7, 8, 9, 11, 12},
			want: 1,
		},
		{
			nums: []int{3, 4, 5, 2, 1},
			want: 6,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFirstMissingPositive(t *testing.T) {
	testFramework(t, firstMissingPositive)
}

func TestFirstMissingPositive_swap(t *testing.T) {
	testFramework(t, firstMissingPositive_swap)
}
