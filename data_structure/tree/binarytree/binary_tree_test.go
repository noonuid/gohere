package binarytree

import (
	"reflect"
	"testing"
)

func TestIntsToTree(t *testing.T) {
	var testCases = []struct {
		ints []int
		want []int
	}{
		{
			ints: []int{},
			want: []int{},
		},
		{
			ints: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			ints: []int{1, Null, 2, 3},
			want: []int{1, 2, 3},
		},
		{
			ints: []int{3, 9, 20, Null, Null, 15, 7},
			want: []int{3, 9, 20, 15, 7},
		},
	}

	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := LevelOrderTraversal(IntsToTree(testCase.ints))

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestTreeToInts(t *testing.T) {
	var testCases = []struct {
		root *TreeNode
		want []int
	}{
		{
			root: IntsToTree([]int{}),
			want: []int{},
		},
		{
			root: IntsToTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			root: IntsToTree([]int{1, Null, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			root: IntsToTree([]int{3, 9, 20, Null, Null, 15, 7}),
			want: []int{3, 9, 20, 15, 7},
		},
	}

	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := TreeToInts(testCase.root)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestLevelOrderTraversal(t *testing.T) {
	var testCases = []struct {
		root *TreeNode
		want []int
	}{
		{
			root: IntsToTree([]int{}),
			want: []int{},
		},
		{
			root: IntsToTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			root: IntsToTree([]int{1, Null, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			root: IntsToTree([]int{3, 9, 20, Null, Null, 15, 7}),
			want: []int{3, 9, 20, 15, 7},
		},
	}

	for caseIndex, testCase := range testCases {
		// 方法的返回值
		got := LevelOrderTraversal(testCase.root)

		// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}

func TestPreorder(t *testing.T) {
	var testCases = []struct {
		root *TreeNode
		want []int
	}{
		{
			root: IntsToTree([]int{}),
			want: []int{},
		},
		{
			root: IntsToTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{1, 2, 4, 5, 3, 6, 7},
		},
		{
			root: IntsToTree([]int{1, Null, 2, 3}),
			want: []int{1, 2, 3},
		},
		{
			root: IntsToTree([]int{3, 9, 20, Null, Null, 15, 7}),
			want: []int{3, 9, 20, 15, 7},
		},
	}

	t.Run("PreorderTraversal", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := PreorderTraversal(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
	t.Run("PreorderTraversalIterative", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := PreorderTraversalIterative(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
}

func TestInorder(t *testing.T) {
	var testCases = []struct {
		root *TreeNode
		want []int
	}{
		{
			root: IntsToTree([]int{}),
			want: []int{},
		},
		{
			root: IntsToTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 2, 5, 1, 6, 3, 7},
		},
		{
			root: IntsToTree([]int{1, Null, 2, 3}),
			want: []int{1, 3, 2},
		},
		{
			root: IntsToTree([]int{3, 9, 20, Null, Null, 15, 7}),
			want: []int{9, 3, 15, 20, 7},
		},
	}

	t.Run("InorderTraversal", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := InorderTraversal(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
	t.Run("InorderTraversalIterative", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := InorderTraversalIterative(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
}

func TestPostorder(t *testing.T) {
	var testCases = []struct {
		root *TreeNode
		want []int
	}{
		{
			root: IntsToTree([]int{}),
			want: []int{},
		},
		{
			root: IntsToTree([]int{1, 2, 3, 4, 5, 6, 7}),
			want: []int{4, 5, 2, 6, 7, 3, 1},
		},
		{
			root: IntsToTree([]int{1, Null, 2, 3}),
			want: []int{3, 2, 1},
		},
		{
			root: IntsToTree([]int{3, 9, 20, Null, Null, 15, 7}),
			want: []int{9, 15, 7, 20, 3},
		},
	}

	t.Run("PostorderTraversal", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := PostorderTraversal(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
	t.Run("PostorderTraversalIterative", func(t *testing.T) {
		for caseIndex, testCase := range testCases {
			// 方法的返回值
			got := PostorderTraversalIterative(testCase.root)

			// 如果方法返回的实际值与期望值不相等，则输出错误对应的测试用例信息
			if !reflect.DeepEqual(got, testCase.want) {
				t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
					caseIndex, testCase, got, testCase.want)
			}
		}
	})
}
