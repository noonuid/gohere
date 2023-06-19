package problem0146

// 请你设计并实现一个满足  LRU (最近最少使用) 缓存 约束的数据结构。

// 实现 LRUCache 类：
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。

// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

type LRUCache struct {
	capacity int
	m        map[int]*Node
	head     *Node
	tail     *Node
}

func (cache *LRUCache) updateNodeOrder(node *Node) {
	if node != cache.head {
		if node.prev != nil {
			node.prev.next = node.next
		}
		if node.next != nil {
			node.next.prev = node.prev
		}
		if node == cache.tail {
			cache.tail = node.prev
		}
		node.prev = nil
		node.next = cache.head
		if cache.head != nil {
			cache.head.prev = node
		}
		cache.head = node
	}
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		m:        make(map[int]*Node, capacity),
	}
}

func (cache *LRUCache) Get(key int) int {
	if node, ok := cache.m[key]; ok && node != nil {
		cache.updateNodeOrder(node)
		return node.value
	} else {
		return -1
	}
}

func (cache *LRUCache) Put(key int, value int) {
	if node, ok := cache.m[key]; ok && node != nil {
		node.value = value
		cache.updateNodeOrder(node)
	} else {
		if len(cache.m) == cache.capacity && cache.tail != nil {
			delete(cache.m, cache.tail.key)
			cache.tail = cache.tail.prev
			if cache.tail != nil {
				cache.tail.next = nil
			}
		}
		node := &Node{
			key:   key,
			value: value,
		}
		if len(cache.m) == 0 {
			cache.head, cache.tail = node, node
		}
		cache.m[key] = node
		cache.updateNodeOrder(node)
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
