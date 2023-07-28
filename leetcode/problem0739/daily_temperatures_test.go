package problem0739

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func([]int) []int) {
	// 测试用例
	testCases := []struct {
		temperatures []int
		want         []int
	}{
		{
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			want:         []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			temperatures: []int{30, 40, 50, 60},
			want:         []int{1, 1, 1, 0},
		},
		{
			temperatures: []int{30, 60, 90},
			want:         []int{1, 1, 0},
		},
		{
			temperatures: []int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70},
			want:         []int{8, 1, 5, 4, 3, 2, 1, 1, 0, 0},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.temperatures)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestDailyTemperatures_stack(t *testing.T) {
	testFramework(t, dailyTemperatures_stack)
}

func TestDailyTemperatures_next(t *testing.T) {
	testFramework(t, dailyTemperatures_next)
}

func TestDailyTemperatures_brute(t *testing.T) {
	testFramework(t, dailyTemperatures_brute)
}
