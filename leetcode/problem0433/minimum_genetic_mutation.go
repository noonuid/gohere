package problem0433

func minMutation(startGene string, endGene string, bank []string) int {
	if startGene == endGene {
		return 0
	}
	diffOne := func(gene1, gene2 string) bool {
		len1, len2 := len(gene1), len(gene2)
		one := false
		for i := 0; i < len1 && i < len2; i++ {
			if gene1[i] != gene2[i] {
				if one {
					return false
				}
				one = true
			}
		}
		return one
	}
	visited := make([]bool, len(bank))
	step := 0
	queue := []string{startGene}
	for len(queue) > 0 {
		step++
		nextQueue := []string{}
		for _, curGene := range queue {
			for index, changedGene := range bank {
				if !visited[index] && diffOne(curGene, changedGene) {
					if changedGene == endGene {
						return step
					}
					visited[index] = true
					nextQueue = append(nextQueue, changedGene)
				}
			}
		}
		queue = nextQueue
	}
	return -1
}
