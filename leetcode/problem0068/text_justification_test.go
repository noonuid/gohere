package problem0068

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(words []string, maxWidth int) []string) {
	testCases := []struct {
		words    []string
		maxWidth int
		want     []string
	}{
		{
			words:    []string{"This", "is", "an", "example", "of", "text", "justification."},
			maxWidth: 16,
			want: []string{
				"This    is    an",
				"example  of text",
				"justification.  ",
			},
		},
		{
			words:    []string{"What", "must", "be", "acknowledgment", "shall", "be"},
			maxWidth: 16,
			want: []string{
				"What   must   be",
				"acknowledgment  ",
				"shall be        ",
			},
		},
		{
			words:    []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			maxWidth: 20,
			want: []string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.words, testCase.maxWidth)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFullJustify(t *testing.T) {
	test(t, fullJustify)
}
