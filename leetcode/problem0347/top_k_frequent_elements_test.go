package problem0347

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) []int) {
	// 测试用例
	testCases := []struct {
		nums []int
		k    int
		want []int
	}{
		{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
			want: []int{1, 2},
		},
		{
			nums: []int{1},
			k:    1,
			want: []int{1},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.nums, testCase.k)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestTopKFrequent(t *testing.T) {
	testFramework(t, topKFrequent)
}
