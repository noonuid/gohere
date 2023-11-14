package problem0416

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) bool) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 5, 11, 5},
			want: true,
		},
		{
			nums: []int{1, 2, 3, 5},
			want: false,
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

func TestCanPartition_dynamic_programming(t *testing.T) {
	testFramework(t, canPartition_dynamic_programming)
}

func TestCanPartition_brute(t *testing.T) {
	testFramework(t, canPartition_brute)
}
