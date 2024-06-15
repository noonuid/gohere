package problem0134

// 134. 加油站

// 在一条环路上有 n 个加油站，其中第 i 个加油站有汽油 gas[i] 升。

// 你有一辆油箱容量无限的的汽车，从第 i 个加油站开往第 i+1 个加油站需要消耗汽油 cost[i] 升。你从其中的一个加油站出发，开始时油箱为空。

// 给定两个整数数组 gas 和 cost ，如果你可以按顺序绕环路行驶一周，则返回出发时加油站的编号，否则返回 -1 。如果存在解，则 保证 它是 唯一 的。

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for start := 0; start < n; {
		sumGas, sumCost, count := 0, 0, 0
		for ; count < n; count++ {
			cur := (start + count) % n
			sumGas = sumGas + gas[cur]
			sumCost = sumCost + cost[cur]
			if sumGas < sumCost {
				break
			}
		}
		if count == n {
			return start
		} else {
			start = start + count + 1
		}
	}
	return -1
}

// 枚举。
// 超出时间限制。
func canCompleteCircuit_enum(gas []int, cost []int) int {
	n := len(gas)
	for start := 0; start < n; start++ {
		remain := 0
		for i := 0; i < n; i++ {
			cur := (start + i) % n
			remain = remain + gas[cur] - cost[cur]
			if remain < 0 {
				break
			}
		}
		if remain >= 0 {
			return start
		}
	}
	return -1
}
