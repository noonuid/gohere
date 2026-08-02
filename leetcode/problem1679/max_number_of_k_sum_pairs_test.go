package problem1679

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
			nums: []int{1, 2, 3, 4},
			k:    5,
			want: 2,
		},
		{
			nums: []int{3, 1, 3, 4, 3},
			k:    6,
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums, testCase.k)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMaxOperations(t *testing.T) {
	testFramework(t, maxOperations)
}

func TestMaxOperations_map(t *testing.T) {
	testFramework(t, maxOperations_map)
}
