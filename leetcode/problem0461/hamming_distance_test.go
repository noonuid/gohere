package problem0461

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(int, int) int) {
	// 测试用例。
	testCases := []struct {
		x    int
		y    int
		want int
	}{
		{
			x:    1,
			y:    4,
			want: 2,
		},
		{
			x:    3,
			y:    1,
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.x, testCase.y)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestHammingDistance_shift_and_count(t *testing.T) {
	testFramework(t, hammingDistance_shift_and_count)
}

func TestHammingDistance(t *testing.T) {
	testFramework(t, hammingDistance_divide_by_two)
}

func TestHammingDistance_ones_count(t *testing.T) {
	testFramework(t, hammingDistance_ones_count)
}
