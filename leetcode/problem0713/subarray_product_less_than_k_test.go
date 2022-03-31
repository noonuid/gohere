package problem0713

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	nums []int
	k    int
	want int
}{
	{
		nums: []int{10, 5, 2, 6},
		k:    100,
		want: 8,
	},
	{
		nums: []int{1, 2, 3},
		k:    0,
		want: 0,
	},
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := numSubarrayProductLessThanK(testCase.nums, testCase.k)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
