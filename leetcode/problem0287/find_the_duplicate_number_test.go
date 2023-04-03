package problem0287

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) int) {
	// 测试用例
	testCases := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			nums: []int{2, 2, 2, 2, 2},
			want: 2,
		},
		{
			nums: []int{1, 4, 6, 6, 6, 2, 3},
			want: 6,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindDuplicate(t *testing.T) {
	testFramework(t, findDuplicate)
}

func TestFindDuplicateBinarySearch(t *testing.T) {
	testFramework(t, findDuplicateBinarySearch)
}

func TestFindFindDuplicateEnum(t *testing.T) {
	testFramework(t, findDuplicateEnum)
}

func TestFindFindDuplicateMap(t *testing.T) {
	testFramework(t, findDuplicateMap)
}
