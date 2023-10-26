package problem0221

import (
	"reflect"
	"testing"
)

func testFramework(t *testing.T, testFunc func(matrix [][]byte) int) {
	// 测试用例。
	testCases := []struct {
		matrix [][]byte
		want   int
	}{
		{
			matrix: [][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}},
			want:   4,
		},
		{
			matrix: [][]byte{{'0', '1'}, {'1', '0'}},
			want:   1,
		},
		{
			matrix: [][]byte{{'0'}},
			want:   0,
		},
	}

	for caseIndex, testCase := range testCases {
		// 被测方法的返回值。
		got := testFunc(testCase.matrix)
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMaximalSquare_dynamic_programming(t *testing.T) {
	testFramework(t, maximalSquare_dynamic_programming)
}
