package nc081

import (
	"reflect"
	"testing"

	"github.com/anchorportal/go/nowcoder/structure"
)

func TestNc081(t *testing.T) {
	testCases := []struct {
		root     []int
		k        int
		expected int
	}{
		{
			root:     []int{},
			k:        1,
			expected: -1,
		},
		{
			root:     []int{3, 1, 4, structure.NULL, 2},
			k:        0,
			expected: -1,
		},
		{
			root:     []int{3, 1, 4, structure.NULL, 2},
			k:        1,
			expected: 1,
		},
		{
			root:     []int{5, 3, 6, 2, 4, structure.NULL, structure.NULL, 1},
			k:        3,
			expected: 3,
		},
	}

	t.Run("KthNode", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := KthNode(structure.Ints2TreeNode(testCase.root), testCase.k)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
