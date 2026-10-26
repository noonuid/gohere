package problem0433

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(startGene string, endGene string, bank []string) int) {
	testCases := []struct {
		startGene string
		endGene   string
		bank      []string
		want      int
	}{
		{
			startGene: "AACCGGTT",
			endGene:   "AACCGGTA",
			bank:      []string{"AACCGGTA"},
			want:      1,
		},
		{
			startGene: "AACCGGTT",
			endGene:   "AAACGGTA",
			bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			want:      2,
		},
		{
			startGene: "AAAAACCC",
			endGene:   "AACCCCCC",
			bank:      []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			want:      3,
		},
		{
			startGene: "AAAAAAAA",
			endGene:   "CCCCCCCC",
			bank:      []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA", "CCCCCCCC"},
			want:      8,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.startGene, testCase.endGene, testCase.bank)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestMinMutation(t *testing.T) {
	test(t, minMutation)
}
