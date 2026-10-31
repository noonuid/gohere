package problem0212

import (
	"reflect"
	"slices"
	"testing"
)

func test(t *testing.T, fn func(board [][]byte, words []string) []string) {
	testCases := []struct {
		board [][]byte
		words []string
		want  []string
	}{
		{
			board: [][]byte{{'o', 'a', 'a', 'n'}, {'e', 't', 'a', 'e'}, {'i', 'h', 'k', 'r'}, {'i', 'f', 'l', 'v'}},
			words: []string{"oath", "pea", "eat", "rain"},
			want:  []string{"eat", "oath"},
		},
		{
			board: [][]byte{{'a', 'b'}, {'c', 'd'}},
			words: []string{"abcb"},
			want:  []string{},
		},
		{
			board: [][]byte{{'a', 'a'}},
			words: []string{"aaa"},
			want:  []string{},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.board, testCase.words)
		slices.Sort(got)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFindWords(t *testing.T) {
	test(t, findWords)
}

func TestFindWords_delete_matched_words(t *testing.T) {
	test(t, findWords_delete_matched_words)
}
