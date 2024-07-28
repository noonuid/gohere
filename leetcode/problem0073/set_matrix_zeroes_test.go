package problem0073

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(matrix [][]int)) {
	testCases := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			matrix: [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}},
			want:   [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			matrix: [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}},
			want:   [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}

	for caseIndex, testCase := range testCases {
		fn(testCase.matrix)
		if !reflect.DeepEqual(testCase.matrix, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, testCase.matrix, testCase.want)
		}
	}
}

func TestSetZeroes(t *testing.T) {
	test(t, setZeroes)
}

func TestSetZeroes_enum(t *testing.T) {
	test(t, setZeroes_enum)
}
