package problem0230

import (
	"reflect"
	"testing"

	"github.com/anchorportal/go/leetcode/structure"
)

func TestProblem0230(t *testing.T) {
	testCases := []struct {
		root     []int
		k        int
		expected int
	}{
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

	t.Run("kthSmallest", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			actual := kthSmallest(structure.Ints2TreeNode(testCase.root), testCase.k)

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\nexpected: %d, actual: %d",
					caseIndex, testCase, testCase.expected, actual)
			}
		}
	})
}
