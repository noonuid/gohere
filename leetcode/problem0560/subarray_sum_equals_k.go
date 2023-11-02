package problem0560

// 560. 和为 K 的子数组

// 给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

// 子数组是数组中元素的连续非空序列。

// 使用哈希表存储前缀和。
func subarraySum_prefix_sum(nums []int, k int) int {
	n, count, prefixSum := len(nums), 0, 0
	// sumCount 记录某个前缀和出现的次数。
	sumCount := make(map[int]int, n+1)
	sumCount[0] = 1
	for i := 0; i < n; i++ {
		prefixSum += nums[i]
		if value, ok := sumCount[prefixSum-k]; ok {
			count += value
		}
		sumCount[prefixSum]++
	}
	return count
}

// 暴力枚举。
func subarraySum_brute(nums []int, k int) int {
	n := len(nums)
	count := 0
	for left := 0; left < n; left++ {
		sum := 0
		for right := left; right < n; right++ {
			sum += nums[right]
			if sum == k {
				count++
			}
		}
	}
	return count
}
