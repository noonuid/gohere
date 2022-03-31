package hj077

import (
	"fmt"
	"sort"
)

func Do() {
	var n, scanN int
	var trains []int

	for {
		scanN, _ = fmt.Scan(&n)
		if scanN == 0 {
			break
		}
		trains = make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&trains[i])
		}
	}

	var sequences = [][]int{}
	trainsPullsInStaion(&sequences, []int{}, []int{}, trains)
	var result = sortResult(sequences)

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

// 回溯法
// sequences 为结果序列，toAdd 为将要添加到 sequences 的不完全出站结果
// stack 存放进站的火车，trains 为火车入站的队列
func trainsPullsInStaion(sequences *[][]int, toAdd, stack, trains []int) {
	// 已完成一次全部火车进站与出战
	if len(stack) == 0 && len(trains) == 0 {
		*sequences = append(*sequences, toAdd)
		return
	}

	// 火车入站的队列不为空，可以选择是否进站
	if len(trains) > 0 {
		// 选择入站，生成新的 stack, trains
		var newStack = make([]int, len(stack), len(stack)+1)
		copy(newStack, stack)
		newStack = append(newStack, trains[0])
		var newTrains = append([]int{}, trains[1:]...)
		trainsPullsInStaion(sequences, toAdd, newStack, newTrains)
		// 递归返回之后，回退选择，继续使用之前的 stack, trains
	}

	// 火车已入站的栈不为空，可以出站，使用 toAdd 存储序列
	if len(stack) > 0 {
		// 生成新的 toAdd, stack
		var newToAdd = make([]int, len(toAdd), len(toAdd)+1)
		copy(newToAdd, toAdd)
		newToAdd = append(newToAdd, stack[len(stack)-1])
		var newStack = append([]int{}, stack[:len(stack)-1]...)
		trainsPullsInStaion(sequences, newToAdd, newStack, trains)
	}
}

// 对出站序列进行排序
func sortResult(sequences [][]int) []string {
	var result = make([]string, len(sequences))

	for i := 0; i < len(result); i++ {
		var str = fmt.Sprint(sequences[i])
		result[i] = str[1 : len(str)-1]
	}

	sort.Strings(result)

	return result
}
