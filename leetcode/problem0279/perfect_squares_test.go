package problem0279

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
			n:    12,
			want: 3,
		},
		{
			n:    13,
			want: 2,
		},
	}

	for caseIndex, testCase := range testCases {
		got := testFunc(testCase.n)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ntest case: %d\ngot : %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestNumSquares(t *testing.T) {
	testFramework(t, numSquares)
}
