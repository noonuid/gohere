package dynamicprograming

// 给定钞票面额分别是 1、5、11 的纸币，
// 求解凑出 total 金额所需的最少钞票数量

// 暴力求解
func minNumOfNotesForce(total int) int {
	var minNum = total
	for i := 0; i < total+1; i++ {
		for j := 0; j < total/5+1; j++ {
			for k := 0; k < total/11+1; k++ {
				if i+j*5+k*11 == total && i+j+k < minNum {
					minNum = i + j + k
				}
			}
		}
	}

	return minNum
}

// 动态规划
func minNumOfNotesDynamicProgramming(total int) int {
	// 状态数组，f[i] 表示凑出 i 所需的最少钞票数量
	var f = make([]int, total+1)
	// 边界为 0
	f[0] = 0

	for cur := 1; cur <= total; cur++ {
		f[cur] = f[cur-1] + 1

		if cur-5 >= 0 && f[cur-5]+1 < f[cur] {
			f[cur] = f[cur-5] + 1
		}
		if cur-11 >= 0 && f[cur-11]+1 < f[cur] {
			f[cur] = f[cur-11] + 1
		}
	}

	return f[total]
}
