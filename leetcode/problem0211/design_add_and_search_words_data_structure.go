package problem0211

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	children    [26]*TrieNode
	isEndOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{root: &TrieNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		childIndex := word[i] - 'a'
		child := node.children[childIndex]
		if child == nil {
			child = &TrieNode{}
			node.children[childIndex] = child
		}
		node = child
	}
	node.isEndOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	n := len(word)

	var search func(node *TrieNode, index int) bool
	search = func(node *TrieNode, index int) bool {
		if index == n {
			return node.isEndOfWord
		}

		if ch := word[index]; ch != '.' {
			childIndex := ch - 'a'
			child := node.children[childIndex]
			return (child != nil) && search(child, index+1)
		} else {
			for _, child := range node.children {
				if (child != nil) && search(child, index+1) {
					return true
				}
			}
			return false
		}
	}

	return search(this.root, 0)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
