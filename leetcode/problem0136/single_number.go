package problem0136

func singleNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		var cur, amount = nums[i], 0
		for j := 0; j < len(nums); j++ {
			if nums[j] == cur {
				amount++
			}
		}
		if amount == 1 {
			return cur
		}
	}
	return -1
}

func singleNumberXor(nums []int) int {
	var result = 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
