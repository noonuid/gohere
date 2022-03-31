package binarysearchtree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)

	if root.data != 90 {
		t.Errorf("root should have value = 90")
	}

	if root.left.data != 80 {
		t.Errorf("left child should have value = 80")
	}

	if root.right.data != 100 {
		t.Errorf("right child should have value = 100")
	}

	if Depth(&tree) != 2 {
		t.Errorf("tree should have depth = 2")
	}
}

func TestPreOrder(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := PreOrder(root)
	b := []int{90, 80, 70, 85, 100, 95, 105}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [90 80 70 85 100 95 105]")
	}
}

func TestInOrder(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := InOrder(root)
	b := []int{70, 80, 85, 90, 95, 100, 105}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [70 80 85 90 95 100 105]")
	}
}

func TestPostOrder(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	a := PostOrder(root)
	b := []int{70, 85, 80, 95, 105, 100, 90}

	if !reflect.DeepEqual(a, b) {
		t.Errorf("Nodes should have value = [70 85 80 95 105 100 90]")
	}
}

func TestMinNode(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	want := 70
	got := MinNode(root).data

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestMaxNode(t *testing.T) {
	tree := BinarySearchTree{
		root: NewNode(90),
	}

	root := tree.root

	Insert(root, 80)
	Insert(root, 100)
	Insert(root, 70)
	Insert(root, 85)
	Insert(root, 95)
	Insert(root, 105)

	want := 105
	got := MaxNode(root).data

	if !reflect.DeepEqual(want, got) {
		t.Errorf("got: %v, want: %v", got, want)
	}
}

func TestRemove(t *testing.T) {
	t.Run("Delete a nil tree", func(t *testing.T) {
		tree := BinarySearchTree{}

		root := tree.root

		var want *Node = nil
		got := Remove(root, 100)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Delete a node with no child", func(t *testing.T) {
		tree := BinarySearchTree{
			root: NewNode(90),
		}

		root := tree.root

		Insert(root, 80)
		Insert(root, 100)

		Remove(root, 100)

		if root.data != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.left.data != 80 {
			t.Errorf("left child should have value = 80")
		}

		if root.right != nil {
			t.Errorf("right child should have value = nil")
		}

		if Depth(&tree) != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with one child", func(t *testing.T) {
		tree := BinarySearchTree{
			root: NewNode(90),
		}

		root := tree.root

		Insert(root, 80)
		Insert(root, 100)
		Insert(root, 70)

		Remove(root, 80)

		if root.data != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.right.data != 100 {
			t.Errorf("right child should have value = 100")
		}

		if root.left.data != 70 {
			t.Errorf("left child should have value = 70")
		}

		if Depth(&tree) != 2 {
			t.Errorf("Depth should have value = 2")
		}
	})

	t.Run("Delete a node with two children", func(t *testing.T) {
		tree := BinarySearchTree{
			root: NewNode(90),
		}

		root := tree.root

		Insert(root, 80)
		Insert(root, 100)
		Insert(root, 70)
		Insert(root, 85)

		Remove(root, 80)

		if root.data != 90 {
			t.Errorf("root should have value = 90")
		}

		if root.left.data != 85 {
			t.Errorf("left child should have value = 85")
		}

		if root.right.data != 100 {
			t.Errorf("right child should have value = 100")
		}

		if Depth(&tree) != 3 {
			t.Errorf("Depth should have value = 3")
		}
	})
}
