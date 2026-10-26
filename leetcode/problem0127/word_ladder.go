package problem0127

import "math"

func ladderLength(beginWord string, endWord string, wordList []string) int {
	diffOne := func(str1, str2 string) bool {
		one := false
		for i := 0; i < len(str1) && i < len(str2); i++ {
			if str1[i] != str2[i] {
				if one {
					return false
				}
				one = true
			}
		}
		return one
	}
	visited := make([]bool, len(wordList))
	length := 1
	queue := []string{beginWord}
	for len(queue) > 0 {
		length++
		nextQueue := []string{}
		for _, curWord := range queue {
			for index, nextWord := range wordList {
				if !visited[index] && diffOne(curWord, nextWord) {
					if nextWord == endWord {
						return length
					}
					visited[index] = true
					nextQueue = append(nextQueue, nextWord)
				}
			}
		}
		queue = nextQueue
	}
	return 0
}

func ladderLength_build_graph(beginWord string, endWord string, wordList []string) int {
	wordIdMap := make(map[string]int)
	edges := make([][]int, 0, len(wordList)*len(beginWord))

	addWord := func(word string) int {
		id, ok := wordIdMap[word]
		if !ok {
			id = len(wordIdMap)
			wordIdMap[word] = id
			edges = append(edges, []int{})
		}
		return id
	}

	addWordToGraph := func(word string) int {
		id1 := addWord(word)
		chs := []byte(word)
		for i, ch := range chs {
			chs[i] = '*'
			id2 := addWord(string(chs))
			edges[id1] = append(edges[id1], id2)
			edges[id2] = append(edges[id2], id1)
			chs[i] = ch
		}
		return id1
	}

	for _, word := range wordList {
		addWordToGraph(word)
	}

	endId, ok := wordIdMap[endWord]
	if !ok {
		return 0
	}

	beginId := addWordToGraph(beginWord)

	const inf = math.MaxInt
	dist := make([]int, len(wordIdMap))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		preQueue := queue
		queue = nil
		for _, curId := range preQueue {
			for _, nextId := range edges[curId] {
				if nextId == endId {
					return (dist[curId]+1)/2 + 1
				}
				if dist[nextId] == inf {
					dist[nextId] = dist[curId] + 1
					queue = append(queue, nextId)
				}
			}
		}
	}
	return 0
}

func ladderLength_bidirectional_bfs(beginWord string, endWord string, wordList []string) int {
	wordIdMap := make(map[string]int)
	edges := make([][]int, 0, len(wordList)*len(beginWord))

	addWord := func(word string) int {
		id, ok := wordIdMap[word]
		if !ok {
			id = len(wordIdMap)
			wordIdMap[word] = id
			edges = append(edges, []int{})
		}
		return id
	}

	addWordToGraph := func(word string) int {
		id1 := addWord(word)
		chs := []byte(word)
		for i, ch := range chs {
			chs[i] = '*'
			id2 := addWord(string(chs))
			edges[id1] = append(edges[id1], id2)
			edges[id2] = append(edges[id2], id1)
			chs[i] = ch
		}
		return id1
	}

	for _, word := range wordList {
		addWordToGraph(word)
	}

	endId, ok := wordIdMap[endWord]
	if !ok {
		return 0
	}

	beginId := addWordToGraph(beginWord)

	const inf = math.MaxInt
	distBegin := make([]int, len(wordIdMap))
	distEnd := make([]int, len(wordIdMap))
	for i := range distBegin {
		distBegin[i] = inf
		distEnd[i] = inf
	}
	distBegin[beginId] = 0
	distEnd[endId] = 0
	queueBegin := []int{beginId}
	queueEnd := []int{endId}
	for len(queueBegin) > 0 && len(queueEnd) > 0 {
		preQueueBegin := queueBegin
		queueBegin = nil
		for _, curId := range preQueueBegin {
			if distEnd[curId] < inf {
				return (distBegin[curId]+distEnd[curId])/2 + 1
			}
			for _, nextId := range edges[curId] {
				if distBegin[nextId] == inf {
					distBegin[nextId] = distBegin[curId] + 1
					queueBegin = append(queueBegin, nextId)
				}
			}
		}

		preQueueEnd := queueEnd
		queueEnd = nil
		for _, curId := range preQueueEnd {
			if distBegin[curId] < inf {
				return (distBegin[curId]+distEnd[curId])/2 + 1
			}
			for _, nextId := range edges[curId] {
				if distEnd[nextId] == inf {
					distEnd[nextId] = distEnd[curId] + 1
					queueEnd = append(queueEnd, nextId)
				}
			}
		}
	}
	return 0
}
