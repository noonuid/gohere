package problem0283

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int)) {
	// 测试用例
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{0, 1, 0, 3, 12},
			want: []int{1, 3, 12, 0, 0},
		},
		{
			nums: []int{0},
			want: []int{0},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(testCase.nums, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, testCase.nums, testCase.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testFramework(t, moveZeroes)
}
