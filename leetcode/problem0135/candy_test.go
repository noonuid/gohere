package problem0135

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(ratings []int) int) {
	testCases := []struct {
		ratings []int
		want    int
	}{
		{
			ratings: []int{1, 0, 2},
			want:    5,
		},
		{
			ratings: []int{1, 2, 2},
			want:    4,
		},
		{
			ratings: []int{1, 3, 5, 2, 3, 3},
			want:    10,
		},
		{
			ratings: []int{1, 3, 5, 3, 2, 2},
			want:    10,
		},
		{
			ratings: []int{1, 3, 5, 3, 2, 1},
			want:    13,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.ratings)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestCandy(t *testing.T) {
	test(t, candy)
}

func TestCandy_one_traversal(t *testing.T) {
	test(t, candy_one_traversal)
}
