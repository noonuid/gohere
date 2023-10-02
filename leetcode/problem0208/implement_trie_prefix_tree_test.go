package problem0208

import (
	"reflect"
	"testing"
)

var null = false

func testFramework(t *testing.T) {
	// 测试用例。
	testCases := []struct {
		operations []string
		paras      [][]string
		want       []bool
	}{
		{
			operations: []string{"Trie", "insert", "search", "search", "startsWith", "insert", "search"},
			paras:      [][]string{{}, {"apple"}, {"apple"}, {"app"}, {"app"}, {"app"}, {"app"}},
			want:       []bool{null, null, true, false, true, null, true},
		},
	}

	for caseIndex, testCase := range testCases {
		var trie Trie
		for i, operation := range testCase.operations {
			var got bool
			switch operation {
			case "Trie":
				trie = Constructor()
			case "insert":
				trie.Insert(testCase.paras[i][0])
			case "search":
				got = trie.Search(testCase.paras[i][0])
			case "startsWith":
				got = trie.StartsWith(testCase.paras[i][0])
			}
			if !reflect.DeepEqual(got, testCase.want[i]) {
				t.Errorf("\ncaseIndex: %d\noperationIndex: %d\ngot: %v\nwant: %v",
					caseIndex, i, got, testCase.want[i])
			}
		}
	}
}

func TestTrie(t *testing.T) {
	testFramework(t)
}
