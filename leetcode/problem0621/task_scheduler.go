package problem0621

// 使用任务和冷却填充 max*(n+1) 的二维矩阵
func leastInterval(tasks []byte, n int) int {
	var freqs = make([]int, 26)
	// 统计每个任务的频率并记录最大值
	for i := 0; i < len(tasks); i++ {
		freqs[tasks[i]-'A']++
	}

	// 找到各个任务频率的最大值
	var maxFreq = 0
	for i := 0; i < 26; i++ {
		if freqs[i] > maxFreq {
			maxFreq = freqs[i]
		}
	}

	var minTime = (maxFreq - 1) * (n + 1)
	for i := 0; i < 26; i++ {
		// 当二维矩阵最后一行中存在频率最大的任务，时间需要增加
		if freqs[i] == maxFreq {
			minTime++
		}
	}

	if minTime > len(tasks) {
		return minTime
	} else {
		return len(tasks)
	}
}
