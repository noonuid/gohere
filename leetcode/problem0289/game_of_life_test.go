package problem0289

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(board [][]int)) {
	testCases := []struct {
		board [][]int
		want  [][]int
	}{
		{
			board: [][]int{
				{0, 1, 0},
				{0, 0, 1},
				{1, 1, 1},
				{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
				{1, 0, 1},
				{0, 1, 1},
				{0, 1, 0},
			},
		},
		{
			board: [][]int{
				{1, 1},
				{1, 0},
			},
			want: [][]int{
				{1, 1},
				{1, 1},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		fn(testCase.board)

		if !reflect.DeepEqual(testCase.want, testCase.board) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, testCase.board)
		}
	}
}

func TestGameOfLife(t *testing.T) {
	test(t, gameOfLife)
}

func TestGameOfLife_composite_state(t *testing.T) {
	test(t, gameOfLife_composite_state)
}
