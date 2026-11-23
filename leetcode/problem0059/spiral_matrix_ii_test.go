package problem0059

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(n int) [][]int) {
	testCases := []struct {
		n    int
		want [][]int
	}{
		{
			n: 3,
			want: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
		{
			n: 1,
			want: [][]int{
				{1},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.n)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestGenerateMatrix(t *testing.T) {
	test(t, generateMatrix)
}

func TestGenerateMatrix_level(t *testing.T) {
	test(t, generateMatrix_level)
}
