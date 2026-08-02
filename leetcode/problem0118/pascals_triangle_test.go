package problem0118

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(numRows int) [][]int) {
	// 测试用例。
	testCases := []struct {
		numRows int
		want    [][]int
	}{
		{
			numRows: 5,
			want:    [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		},
		{
			numRows: 1,
			want:    [][]int{{1}},
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.numRows)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestGenerate(t *testing.T) {
	testFramework(t, generate)
}
