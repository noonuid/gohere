package structure

import (
	"reflect"
	"testing"
)

func TestInts2Tree(t *testing.T) {
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
			nums:      []int{1, 2, NULL, 3},
			preorder:  []int{1, 2, 3},
			inorder:   []int{3, 2, 1},
			postorder: []int{3, 2, 1},
		},
		{
			nums:      []int{1, NULL, 2, NULL, 3},
			preorder:  []int{1, 2, 3},
			inorder:   []int{1, 2, 3},
			postorder: []int{3, 2, 1},
		},
		{
			nums:      []int{1, 2, 3, 4, NULL, NULL, 7},
			preorder:  []int{1, 2, 4, 3, 7},
			inorder:   []int{4, 2, 1, 3, 7},
			postorder: []int{4, 2, 7, 3, 1},
		},
	}

	for caseIndex, testCase := range testCases {
		root := Ints2Tree(testCase.nums)

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

func TestTree2Ints(t *testing.T) {
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
			nums: []int{1, 2, NULL, 3},
			want: []int{1, 2, NULL, 3},
		},
		{
			nums: []int{1, NULL, 2, NULL, 3},
			want: []int{1, NULL, 2, NULL, 3},
		},
		{
			nums: []int{1, 2, 3, 4, NULL, NULL, 7},
			want: []int{1, 2, 3, 4, NULL, NULL, 7},
		},
		{
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for caseIndex, testCase := range testCases {
		root := Ints2Tree(testCase.nums)
		got := Tree2Ints(root)
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("\ncaseIndex: %d, testCase: %v\ngot: %v, want: %v",
				caseIndex, testCase, got, testCase.want)
		}
	}
}
