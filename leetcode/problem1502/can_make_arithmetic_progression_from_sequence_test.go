package problem1502

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(arr []int) bool) {
	testCases := []struct {
		arr  []int
		want bool
	}{
		{
			arr:  []int{3, 5, 1},
			want: true,
		},
		{
			arr:  []int{1, 2, 4},
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.arr)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCanMakeArithmeticProgression(t *testing.T) {
	test(t, canMakeArithmeticProgression)
}
