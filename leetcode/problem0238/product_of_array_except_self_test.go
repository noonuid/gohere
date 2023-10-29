package problem0238

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
			nums: []int{1, 2, 3, 4},
			want: []int{24, 12, 8, 6},
		},
		{
			nums: []int{-1, 1, 0, -3, 3},
			want: []int{0, 0, 9, 0, 0},
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

func TestProductExceptSelf_dynamic_programming_space_optimization(t *testing.T) {
	testFramework(t, productExceptSelf_dynamic_programming_space_optimization)
}

func TestProductExceptSelf_dynamic_programming(t *testing.T) {
	testFramework(t, productExceptSelf_dynamic_programming)
}

func TestProductExceptSelf_brute(t *testing.T) {
	testFramework(t, productExceptSelf_brute)
}
