package problem0448

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) []int) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{4, 3, 2, 7, 8, 2, 3, 1},
			want: []int{5, 6},
		},
		{
			nums: []int{1, 1},
			want: []int{2},
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

func TestFindDisappearedNumbers_nums(t *testing.T) {
	testFramework(t, findDisappearedNumbers_nums)
}

func TestFindDisappearedNumbers_map(t *testing.T) {
	testFramework(t, findDisappearedNumbers_map)
}
