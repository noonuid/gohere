package problem0075

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int)) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{2, 0, 2, 1, 1, 0},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			nums: []int{2, 0, 1},
			want: []int{0, 1, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(testCase.nums, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.nums, testCase.want)
		}
	}
}

func TestSortColors_double_pointer(t *testing.T) {
	testFramework(t, sortColors_double_pointer)
}

func TestSortColors_single_pointer(t *testing.T) {
	testFramework(t, sortColors_single_pointer)
}

func TestSortColors_bubble(t *testing.T) {
	testFramework(t, sortColors_bubble)
}
