package nc078

import (
	"reflect"
	"testing"

	"github.com/anchorportal/go/nowcoder/structure"
)

func TestNc078(t *testing.T) {
	t.Run("ReverseList", func(t *testing.T) {
		testCases := []struct {
			input    []int
			expected []int
		}{
			{
				input:    []int{1, 2, 3},
				expected: []int{3, 2, 1},
			},
			{
				input:    []int{},
				expected: []int{},
			},
		}

		for _, testCase := range testCases {
			actual := structure.List2Ints(ReverseList(structure.Ints2List(testCase.input)))

			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("case: %v, actual: %d, expected: %d", testCase, actual, testCase.expected)
			}
		}
	})
}
