package problem0128

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		{
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		{
			nums: []int{1, 2, 0, 1},
			want: 3,
		},
		{
			nums: []int{0, 0, -1},
			want: 2,
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

func TestLongestConsecutive_map(t *testing.T) {
	testFramework(t, longestConsecutive_map)
}

func TestLongestConsecutive_sort(t *testing.T) {
	testFramework(t, longestConsecutive_sort)
}
