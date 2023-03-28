package problem0033

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int, target int) int) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			nums:   []int{1, 3},
			target: 1,
			want:   0,
		},
		{
			nums:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			target: 5,
			want:   4,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.nums, testCase.target)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestSearch(t *testing.T) {
	testFramework(t, search)
}

func TestSearchMin(t *testing.T) {
	testFramework(t, searchMin)
}

func TestSearchEnum(t *testing.T) {
	testFramework(t, searchEnum)
}
