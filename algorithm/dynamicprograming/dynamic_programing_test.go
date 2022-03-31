package dynamicprograming

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	total int
	want  int
}{
	{
		total: 0,
		want:  0,
	},
	{
		total: 1,
		want:  1,
	},
	{
		total: 5,
		want:  1,
	},
	{
		total: 11,
		want:  1,
	},
	{
		total: 6,
		want:  2,
	},
	{
		total: 15,
		want:  3,
	},
}

func TestMinNumOfNotesForce(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = minNumOfNotesForce(testCase.total)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestMinNumOfNotesDynamicProgramming(t *testing.T) {
	for caseIndex, testCase := range testCases {
		var got = minNumOfNotesDynamicProgramming(testCase.total)

		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
