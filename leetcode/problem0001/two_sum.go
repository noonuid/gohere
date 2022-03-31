package problem0001

// 方法一：暴力枚举
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

// 方法二：借助哈希表存储第一个数
func twoSumMap(nums []int, target int) []int {
	var visited = make(map[int]int, len(nums))

	for second, num := range nums {
		if first, ok := visited[target-num]; ok {
			return []int{first, second}
		} else {
			visited[num] = second
		}
	}

	return nil
}
