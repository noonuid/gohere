package binarytree

import (
	"reflect"
	"testing"
)

func testIntsToTree(t *testing.T, testFn func([]int) *TreeNode) {
	testCases := []struct {
		nums      []int
		preorder  []int
		inorder   []int
		postorder []int
	}{
		{
			nums:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			preorder:  []int{1, 2, 4, 8, 9, 5, 10, 11, 3, 6, 12, 13, 7, 14, 15},
			inorder:   []int{8, 4, 9, 2, 10, 5, 11, 1, 12, 6, 13, 3, 14, 7, 15},
			postorder: []int{8, 9, 4, 10, 11, 5, 2, 12, 13, 6, 14, 15, 7, 3, 1},
		},
		{
			nums:      []int{},
			preorder:  []int{},
			inorder:   []int{},
			postorder: []int{},
		},
		{
			nums:      []int{1},
			preorder:  []int{1},
			inorder:   []int{1},
			postorder: []int{1},
		},
		{
			nums:      []int{1, 2, Null, 3},
			preorder:  []int{1, 2, 3},
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
		},
		{
			nums:      []int{1, Null, 2, Null, Null, Null, 3},
			preorder:  []int{1, 2, 3},
			inorder:   []int{1, 2, 3},
			postorder: []int{3, 2, 1},
		},
		{
			nums:      []int{1, 2, 3, 4, Null, Null, 7},
			preorder:  []int{1, 2, 4, 3, 7},
			inorder:   []int{4, 2, 1, 3, 7},
			postorder: []int{4, 2, 7, 3, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		root := testFn(testCase.nums)

		preorder := PreorderTraversal(root)
		if !reflect.DeepEqual(preorder, testCase.preorder) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot preorder: %v, want: %v",
				caseIndex, testCase, preorder, testCase.preorder)
		}

		inorder := InorderTraversal(root)
		if !reflect.DeepEqual(inorder, testCase.inorder) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot inorder: %v, want: %v",
				caseIndex, testCase, inorder, testCase.inorder)
		}

		postorder := PostorderTraversal(root)
		if !reflect.DeepEqual(postorder, testCase.postorder) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot postorder: %v, want: %v",
				caseIndex, testCase, postorder, testCase.postorder)
		}
	}
}

func TestIntsToTree(t *testing.T) {
	testIntsToTree(t, IntsToTree)
}

func TestIntsToTree_recurse(t *testing.T) {
	testIntsToTree(t, IntsToTree_recurse)
}

func TestTreeToInts(t *testing.T) {
	testCases := []struct {
		nums []int
		want []int
	}{
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7},
			want: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			nums: []int{},
			want: []int{},
		},
		{
			nums: []int{1},
			want: []int{1},
		},
		{
			nums: []int{1, 2, Null, 3},
			want: []int{1, 2, Null, 3},
		},
		{
			nums: []int{1, Null, 2, Null, Null, Null, 3},
			want: []int{1, Null, 2, Null, Null, Null, 3},
		},
		{
			nums: []int{1, 2, 3, 4, Null, Null, 7},
			want: []int{1, 2, 3, 4, Null, Null, 7},
		},
	}

	for caseIndex, testCase := range testCases {
		root := IntsToTree(testCase.nums)
		got := TreeToInts(root)
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
			root: IntsToTree([]int{1, Null, 2, Null, Null, 3}),
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
			root: IntsToTree([]int{1, Null, 2, Null, Null, 3}),
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
			root: IntsToTree([]int{1, Null, 2, Null, Null, 3}),
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
			root: IntsToTree([]int{1, Null, 2, Null, Null, 3}),
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
