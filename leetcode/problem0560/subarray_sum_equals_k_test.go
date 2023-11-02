package problem0560

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int, k int) int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		k    int
		want int
	}{
		{
			nums: []int{1, 1, 1},
			k:    2,
			want: 2,
		},
		{
			nums: []int{1, 2, 3},
			k:    3,
			want: 2,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums, testCase.k)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestSubarraySum_prefix_sum(t *testing.T) {
	testFramework(t, subarraySum_prefix_sum)
}

func TestSubarraySum_brute(t *testing.T) {
	testFramework(t, subarraySum_brute)
}
