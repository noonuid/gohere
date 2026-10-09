package problem0380

import (
	"reflect"
	"slices"
	"testing"
)

func test(t *testing.T, fn func() RandomizedSet) {
	testCases := []struct {
		ops   []string
		vals  [][]int
		wants []any
	}{
		{
			ops: []string{
				"RandomizedSet",
				"insert",
				"remove",
				"insert",
				"getRandom",
				"remove",
				"insert",
				"getRandom",
			},
			vals: [][]int{
				{},
				{1},
				{2},
				{2},
				{},
				{1},
				{2},
				{},
			},
			wants: []any{
				nil,
				true,
				false,
				true,
				[]int{1, 2},
				true,
				false,
				[]int{2},
			},
		},
	}

	for caseIndex, testCase := range testCases {
		var rs RandomizedSet
		gots := make([]any, 0, len(testCase.ops))

		for i, op := range testCase.ops {
			switch op {
			case "RandomizedSet":
				rs = fn()
				gots = append(gots, nil)

			case "insert":
				got := rs.Insert(testCase.vals[i][0])
				gots = append(gots, got)

			case "remove":
				got := rs.Remove(testCase.vals[i][0])
				gots = append(gots, got)

			case "getRandom":
				got := rs.GetRandom()
				gots = append(gots, got)
			}
		}

		for i, op := range testCase.ops {
			switch op {
			case "RandomizedSet":
				if rs.indices == nil || rs.nums == nil {
					t.Errorf("\ncase index: %d\nConstructor failed",
						caseIndex)
				}

			case "insert":
				if !reflect.DeepEqual(testCase.wants[i], gots[i]) {
					t.Errorf("\ncase index: %d\ninsert\nwant: %v\ngot : %v",
						caseIndex, testCase.wants[i], gots[i])
				}

			case "remove":
				if !reflect.DeepEqual(testCase.wants[i], gots[i]) {
					t.Errorf("\ncase index: %d\nremove\nwant: %v\ngot : %v",
						caseIndex, testCase.wants[i], gots[i])
				}

			case "getRandom":
				if !slices.Contains(testCase.wants[i].([]int), gots[i].(int)) {
					t.Errorf("\ncase index: %d\ngetRandom\nwant: %v\ngot : %v",
						caseIndex, testCase.wants[i], gots[i])
				}
			}
		}
	}
}

func TestRandomizedSet(t *testing.T) {
	test(t, Constructor)
}
