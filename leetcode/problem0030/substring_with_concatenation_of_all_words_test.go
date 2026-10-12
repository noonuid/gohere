package problem0030

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(s string, words []string) []int) {
	testCases := []struct {
		s     string
		words []string
		want  []int
	}{
		{
			s:     "barfoothefoobarman",
			words: []string{"foo", "bar"},
			want:  []int{0, 9},
		},
		{
			s:     "wordgoodgoodgoodbestword",
			words: []string{"word", "good", "best", "word"},
			want:  []int{},
		},
		{
			s:     "barfoofoobarthefoobarman",
			words: []string{"bar", "foo", "the"},
			want:  []int{6, 9, 12},
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.s, testCase.words)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestFindSubstring(t *testing.T) {
	test(t, findSubstring)
}

func TestFindSubstring_map(t *testing.T) {
	test(t, findSubstring_map)
}
