package problem0338

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(int) []int) {
	// 测试用例。
	testCases := []struct {
		n    int
		want []int
	}{
		{
			n:    2,
			want: []int{0, 1, 1},
		},
		{
			n:    5,
			want: []int{0, 1, 1, 2, 1, 2},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.n)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCountBits_brian_kernighan(t *testing.T) {
	testFramework(t, countBits_brian_kernighan)
}

func TestCountBits_dynamic_programming(t *testing.T) {
	testFramework(t, countBits_dynamic_programming)
}
