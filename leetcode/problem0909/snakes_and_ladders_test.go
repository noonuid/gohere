package problem0909

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(board [][]int) int) {
	testCases := []struct {
		board [][]int
		want  int
	}{
		{
			board: [][]int{
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 35, -1, -1, 13, -1},
				{-1, -1, -1, -1, -1, -1},
				{-1, 15, -1, -1, -1, -1},
			},
			want: 4,
		},
		{
			board: [][]int{
				{-1, -1},
				{-1, 3},
			},
			want: 1,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.board)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestSnakesAndLadders(t *testing.T) {
	test(t, snakesAndLadders)
}
