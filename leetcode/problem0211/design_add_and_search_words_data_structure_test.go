package problem0211

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func() WordDictionary) {
	testCases := []struct {
		operations []string
		parameters [][]string
		want       []any
	}{
		{
			operations: []string{"WordDictionary", "addWord", "addWord", "addWord", "search", "search", "search", "search"},
			parameters: [][]string{{}, {"bad"}, {"dad"}, {"mad"}, {"pad"}, {"bad"}, {".ad"}, {"b.."}},
			want:       []any{nil, nil, nil, nil, false, true, true, true},
		},
	}

	for caseIndex, testCase := range testCases {
		var wordDictionary WordDictionary
		got := []any{}

		for i, operation := range testCase.operations {
			switch operation {
			case "WordDictionary":
				wordDictionary = fn()
				got = append(got, nil)
			case "addWord":
				wordDictionary.AddWord(testCase.parameters[i][0])
				got = append(got, nil)
			case "search":
				got = append(got, wordDictionary.Search(testCase.parameters[i][0]))
			}
		}

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestWordDictionary(t *testing.T) {
	test(t, Constructor)
}
