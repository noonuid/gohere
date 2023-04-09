package problem0051

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(n int) [][]string) {
	// 测试用例
	testCases := []struct {
		n    int
		want [][]string
	}{
		{
			n: 4,
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."}},
		},
		{
			n:    1,
			want: [][]string{{"Q"}},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值
		got := testFunc(testCase.n)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestSolveNQueens(t *testing.T) {
	testFramework(t, solveNQueens)
}
