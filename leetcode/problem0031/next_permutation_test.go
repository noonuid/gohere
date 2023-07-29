package problem0031

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
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(testCase.nums, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, testCase.nums, testCase.want)
		}
	}
}

func TestNextPermutation(t *testing.T) {
	testFramework(t, nextPermutation)
}
