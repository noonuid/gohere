package problem0138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	var headNew *Node
	if head == nil {
		return headNew
	}
	headNew = &Node{Val: head.Val}
	for node, nodeNew := head.Next, headNew; node != nil; node, nodeNew = node.Next, nodeNew.Next {
		nodeNew.Next = &Node{Val: node.Val}
	}
	for node, nodeNew := head, headNew; node != nil; node, nodeNew = node.Next, nodeNew.Next {
		if node.Random != nil {
			random, randomNew := head, headNew
			for random != node.Random {
				random, randomNew = random.Next, randomNew.Next
			}
			nodeNew.Random = randomNew
		}
	}
	return headNew
}

func copyRandomList_hash_table(head *Node) *Node {
	m := make(map[*Node]*Node)
	for node := head; node != nil; node = node.Next {
		m[node] = &Node{Val: node.Val}
	}
	for node := head; node != nil; node = node.Next {
		m[node].Next = m[node.Next]
		m[node].Random = m[node.Random]
	}
	return m[head]
}

func copyRandomList_interleave_nodes(head *Node) *Node {
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node{Val: node.Val, Next: node.Next}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	headNew := head.Next
	for node := head; node != nil; node = node.Next {
		nodeNew := node.Next
		node.Next = node.Next.Next
		if nodeNew.Next != nil {
			nodeNew.Next = node.Next.Next
		}
	}
	return headNew
}
