package problem0155

import (
	"math"
	"reflect"
	"testing"
)

var null = math.MinInt

func testFramework(t *testing.T) {
	// 测试用例。
	testCases := []struct {
		operations []string
		operands   [][]int
		want       []int
	}{
		{
			operations: []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
			operands:   [][]int{{}, {-2}, {0}, {-3}, {}, {}, {}, {}},
			want:       []int{null, null, null, null, -3, null, 0, -2},
		},
	}

	for caseIndex, testCase := range testCases {
		var obj MinStack
		got := []int{}
		for index, operation := range testCase.operations {
			switch operation {
			case "MinStack":
				obj = Constructor()
				got = append(got, null)
			case "push":
				obj.Push(testCase.operands[index][0])
				got = append(got, null)
			case "pop":
				obj.Pop()
				got = append(got, null)
			case "top":
				top := obj.Top()
				got = append(got, top)
			case "getMin":
				min := obj.GetMin()
				got = append(got, min)
			default:
			}
		}
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d\ngot: %v\nwant: %v",
				caseIndex, got, testCase.want)
		}
	}
}

func TestMinStack(t *testing.T) {
	testFramework(t)
}
