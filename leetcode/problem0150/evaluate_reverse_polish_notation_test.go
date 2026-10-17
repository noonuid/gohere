package problem0150

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(tokens []string) int) {
	testCases := []struct {
		tokens []string
		want   int
	}{
		{
			tokens: []string{"2", "1", "+", "3", "*"},
			want:   9,
		},
		{
			tokens: []string{"4", "13", "5", "/", "+"},
			want:   6,
		},
		{
			tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			want:   22,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.tokens)

		if !reflect.DeepEqual(testCase.want, got) {
			t.Errorf("\ncase index: %d\nwant: %v\ngot : %v",
				caseIndex, testCase.want, got)
		}
	}
}

func TestEvalRPN(t *testing.T) {
	test(t, evalRPN)
}

func TestEvalRPN_array_stack(t *testing.T) {
	test(t, evalRPN_array_stack)
}
