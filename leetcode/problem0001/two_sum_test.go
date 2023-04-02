package problem0001

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int, int) []int) {
	// 测试用例
	testCases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			nums:   []int{3, 2, 4},
			target: 5,
			want:   []int{0, 1},
		},
		{
			nums:   []int{0, 8, 7, 3, 3, 4, 2},
			target: 11,
			want:   []int{1, 3},
		},
		{
			nums:   []int{0, 1},
			target: 1,
			want:   []int{0, 1},
		},
		{
			nums:   []int{0, 3},
			target: 5,
			want:   nil,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.nums, testCase.target)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestTwoSum(t *testing.T) {
	testFramework(t, twoSum)
}

func TestTwoSumMap(t *testing.T) {
	testFramework(t, twoSumMap)
}
