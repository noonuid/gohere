package problem0009

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(x int) bool) {
	testCases := []struct {
		x    int
		want bool
	}{
		{
			x:    121,
			want: true,
		},
		{
			x:    -121,
			want: false,
		},
		{
			x:    10,
			want: false,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.x)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	test(t, isPalindrome)
}
