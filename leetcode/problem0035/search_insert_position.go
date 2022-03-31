package problem0035

func searchInsert(nums []int, target int) int {
	if target <= nums[0] {
		return 0
	} else if target > nums[len(nums)-1] {
		return len(nums)
	}

	low := 0
	high := len(nums) - 1
	delta := 0.0
	position := 0

	for low <= high && target >= nums[low] && target <= nums[high] {
		delta = float64(target-nums[low]) / float64(nums[high]-nums[low])
		position = int(float64(low) + float64(high-low)*delta)

		if nums[position] < target {
			low = position + 1
		} else if nums[position] > target {
			high = position - 1
		} else {
			return position
		}
	}

	if nums[position] < target {
		return position + 1
	} else {
		return position
	}
}
