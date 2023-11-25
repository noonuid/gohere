package problem0062

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(m int, n int) int) {
	// 测试用例。
	testCases := []struct {
		m    int
		n    int
		want int
	}{
		{
			m:    3,
			n:    7,
			want: 28,
		},
		{
			m:    3,
			n:    2,
			want: 3,
		},
		{
			m:    3,
			n:    3,
			want: 6,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.m, testCase.n)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestUniquePaths(t *testing.T) {
	testFramework(t, uniquePaths)
}
