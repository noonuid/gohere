package problem0283

/*
283. 移动零

给定一个数组 `nums`，编写一个函数将所有 `0` 移动到数组的末尾，同时保持非零元素的相对顺序。

**请注意** ，必须在不复制数组的情况下原地对数组进行操作。

**示例 1**：

**输入**： nums = `[0,1,0,3,12]`
**输出**： `[1,3,12,0,0]`

**示例 2**：

**输入**： nums = `[0]`
**输出**： `[0]`

**提示**：

- `1 <= nums.length <= 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

**进阶**：你能尽量减少完成的操作次数吗？
*/

// 快慢指针。
func moveZeroes(nums []int) {
	// slow 指向已遍历元素中最后最后一个非 0 元素。
	slow, fast := -1, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
}
