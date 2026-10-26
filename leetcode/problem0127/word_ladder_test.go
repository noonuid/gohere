package problem0127

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(beginWord string, endWord string, wordList []string) int) {
	testCases := []struct {
		beginWord string
		endWord   string
		wordList  []string
		want      int
	}{
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			want:      5,
		},
		{
			beginWord: "hit",
			endWord:   "cog",
			wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			want:      0,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.beginWord, testCase.endWord, testCase.wordList)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestLadderLength(t *testing.T) {
	test(t, ladderLength)
}

func TestLadderLength_build_graph(t *testing.T) {
	test(t, ladderLength_build_graph)
}

func TestLadderLength_bidirectional_bfs(t *testing.T) {
	test(t, ladderLength_bidirectional_bfs)
}
