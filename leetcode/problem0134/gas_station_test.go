package problem0134

import (
	"reflect"
	"testing"
)

func test(t *testing.T, fn func(gas []int, cost []int) int) {
	testCases := []struct {
		gas  []int
		cost []int
		want int
	}{
		{
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			gas:  []int{5, 1, 2, 3, 4},
			cost: []int{4, 4, 1, 5, 1},
			want: 4,
		},
	}

	for caseIndex, testCase := range testCases {
		got := fn(testCase.gas, testCase.cost)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestCanCompleteCircuit(t *testing.T) {
	test(t, canCompleteCircuit)
}

func TestCanCompleteCircuit_enum(t *testing.T) {
	test(t, canCompleteCircuit_enum)
}
