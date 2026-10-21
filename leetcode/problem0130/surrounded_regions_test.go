package problem0130

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(board [][]byte)) {
	testCases := []struct {
		board [][]byte
		want  [][]byte
	}{
		{
			board: [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}},
			want:  [][]byte{{'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'X', 'X', 'X'}, {'X', 'O', 'X', 'X'}},
		},
		{
			board: [][]byte{{'X'}},
			want:  [][]byte{{'X'}},
		},
	}

	for caseIndex, testCase := range testCases {
		fn(testCase.board)
		got := testCase.board

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestSolve(t *testing.T) {
	test(t, solve)
}
