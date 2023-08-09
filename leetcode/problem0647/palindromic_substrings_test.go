package problem0647

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(s string) int) {
	// 测试用例。
	testCases := []struct {
		s    string
		want int
	}{
		{
			s:    "abc",
			want: 3,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.s)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestCountSubstrings_center_expand(t *testing.T) {
	testFramework(t, countSubstrings_center_expand)
}

func TestCountSubstrings_dynamic_programming(t *testing.T) {
	testFramework(t, countSubstrings_dynamic_programming)
}

func TestCountSubstrings_brute(t *testing.T) {
	testFramework(t, countSubstrings_brute)
}
