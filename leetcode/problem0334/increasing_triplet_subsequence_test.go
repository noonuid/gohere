package problem0334

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(nums []int) bool) {
	// 测试用例。
	testCases := []struct {
		nums []int
		want bool
	}{
		{
			nums: []int{1, 2, 3, 4, 5},
			want: true,
		},
		{
			nums: []int{5, 4, 3, 2, 1},
			want: false,
		},
		{
			nums: []int{2, 1, 5, 0, 4, 6},
			want: true,
		},
		{
			nums: []int{20, 100, 10, 12, 5, 13},
			want: true,
		},
		{
			nums: []int{0, 4, 2, 1, 0, -1, -3},
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.nums)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestIncreasingTriplet(t *testing.T) {
	testFramework(t, increasingTriplet)
}

func TestIncreasingTriplet_greedy(t *testing.T) {
	testFramework(t, increasingTriplet_greedy)
}
