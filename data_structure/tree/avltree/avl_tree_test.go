package avltree

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Run("llRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 5)
		tree.root = Insert(tree.root, 3)
		tree.root = Insert(tree.root, 4)

		want := 4
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 3
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 5
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}

		want = 2
		got = Depth(tree)
		if want != got {
			t.Errorf("Depth: want: %v, got: %v", want, got)
		}
	})

	t.Run("lrRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 5)
		tree.root = Insert(tree.root, 4)
		tree.root = Insert(tree.root, 3)

		want := 4
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 3
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 5
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}

		want = 2
		got = Depth(tree)
		if want != got {
			t.Errorf("Depth: want: %v, got: %v", want, got)
		}
	})

	t.Run("rrRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 3)
		tree.root = Insert(tree.root, 4)
		tree.root = Insert(tree.root, 5)

		if tree.root.data != 4 {
			t.Errorf("root should have value = 4")
		}

		if tree.root.left.data != 3 {
			t.Errorf("left child should have value = 3")
		}

		if tree.root.right.data != 5 {
			t.Errorf("right child should have value = 5")
		}

		want := 2
		got := Depth(tree)

		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})

	t.Run("rlRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 3)
		tree.root = Insert(tree.root, 5)
		tree.root = Insert(tree.root, 4)

		if tree.root.data != 4 {
			t.Errorf("root should have value = 4")
		}

		if tree.root.left.data != 3 {
			t.Errorf("left child should have value = 3")
		}

		if tree.root.right.data != 5 {
			t.Errorf("right child should have value = 5")
		}

		want := 2
		got := Depth(tree)

		if want != got {
			t.Errorf("want: %v, got: %v", want, got)
		}
	})
}

func TestPreOrder(t *testing.T) {
	tree := AvlTree{
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
	tree := AvlTree{
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
	tree := AvlTree{
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
	tree := AvlTree{
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
	tree := AvlTree{
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
		tree := AvlTree{}

		root := tree.root

		var want *Node = nil
		got := Remove(root, 100)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})

	t.Run("Delete a node with no child", func(t *testing.T) {
		tree := AvlTree{
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
		tree := AvlTree{
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
		tree := AvlTree{
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

	t.Run("llRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 5)
		tree.root = Insert(tree.root, 4)
		tree.root = Insert(tree.root, 3)
		tree.root = Insert(tree.root, 2)

		tree.root = Remove(tree.root, 5)

		want := 3
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 2
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 4
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}
	})

	t.Run("lrRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 10)
		tree.root = Insert(tree.root, 8)
		tree.root = Insert(tree.root, 6)
		tree.root = Insert(tree.root, 7)

		tree.root = Remove(tree.root, 10)

		want := 7
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 6
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 8
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}
	})

	t.Run("rrRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 2)
		tree.root = Insert(tree.root, 3)
		tree.root = Insert(tree.root, 4)
		tree.root = Insert(tree.root, 5)

		tree.root = Remove(tree.root, 2)

		want := 4
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 3
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 5
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}
	})

	t.Run("rlRotaion", func(t *testing.T) {
		tree := NewAvlTree()

		tree.root = Insert(tree.root, 7)
		tree.root = Insert(tree.root, 6)
		tree.root = Insert(tree.root, 9)
		tree.root = Insert(tree.root, 8)

		tree.root = Remove(tree.root, 6)

		want := 8
		got := tree.root.data
		if want != got {
			t.Errorf("root: want: %v, got: %v", want, got)
		}

		want = 7
		got = tree.root.left.data
		if want != got {
			t.Errorf("left child: want: %v, got: %v", want, got)
		}

		want = 9
		got = tree.root.right.data
		if want != got {
			t.Errorf("right child: want: %v, got: %v", want, got)
		}
	})
}
