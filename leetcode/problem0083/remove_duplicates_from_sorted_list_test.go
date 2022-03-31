package problem0083

import (
	"fmt"
	"testing"

	"github.com/anchorportal/go/leetcode/structure"
)

type question83 struct {
	para83
	ans83
}

// para 是参数
// one 代表第一个参数
type para83 struct {
	one []int
}

// ans 是答案
// one 代表第一个答案
type ans83 struct {
	one []int
}

func Test_Problem83(t *testing.T) {

	qs := []question83{

		{
			para83{[]int{1, 1, 2}},
			ans83{[]int{1, 2}},
		},

		{
			para83{[]int{1, 1, 2, 2, 3, 3, 3}},
			ans83{[]int{1, 2, 3}},
		},

		{
			para83{[]int{1, 1, 1, 1, 1, 1, 1, 1}},
			ans83{[]int{1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 83------------------------\n")

	for _, q := range qs {
		_, p := q.ans83, q.para83
		fmt.Printf("【input】:%v       【output】:%v\n",
			p, structure.List2Ints(deleteDuplicates(structure.Ints2List(p.one))))
	}
	fmt.Printf("\n\n\n")
}
