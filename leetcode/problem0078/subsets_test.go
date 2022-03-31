package problem0078

import (
	"reflect"
	"testing"
)

// 测试用例
var testCases = []struct {
	nums []int
	want [][]int
}{
	{
		nums: []int{1, 2, 3},
		want: [][]int{
			{},
			{1},
			{2},
			{1, 2},
			{3},
			{1, 3},
			{2, 3},
			{1, 2, 3},
		},
	},
	{
		nums: []int{0},
		want: [][]int{
			{},
			{0},
		},
	},
}

func TestSubsets(t *testing.T) {
	for caseIndex, testCase := range testCases {
		// 方法的返回值
		var got = subsets(testCase.nums)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
