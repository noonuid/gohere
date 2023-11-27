package problem0096

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(n int) int) {
	// 测试用例。
	testCases := []struct {
		n    int
		want int
	}{
		{
			n:    3,
			want: 5,
		},
		{
			n:    1,
			want: 1,
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

func TestNumTrees(t *testing.T) {
	testFramework(t, numTrees)
}
