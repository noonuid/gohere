package problem0295

import (
	"math"
	"reflect"
	"testing"
)

var null float64 = math.Pow(10, 5) + 1

func testFramework(t *testing.T) {
	// 测试用例。
	testCases := []struct {
		funcs []string
		paras [][]int
		want  []float64
	}{
		{
			funcs: []string{"MedianFinder", "addNum", "addNum", "findMedian", "addNum", "findMedian"},
			paras: [][]int{{}, {1}, {2}, {}, {3}, {}},
			want:  []float64{null, null, null, 1.5, null, 2.0},
		},
	}

	for caseIndex, testCase := range testCases {
		var mf MedianFinder
		// 被测方法的返回值。
		got := []float64{}
		for i, fn := range testCase.funcs {
			switch fn {
			case "MedianFinder":
				mf = Constructor()
				got = append(got, null)
			case "addNum":
				mf.AddNum(testCase.paras[i][0])
				got = append(got, null)
			case "findMedian":
				median := mf.FindMedian()
				got = append(got, median)
			}
		}
		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息。
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncase index:\t%d\ntest case:\t%v\ngot:\t\t%v\nwant:\t\t%v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestFindMedian(t *testing.T) {
	testFramework(t)
}
