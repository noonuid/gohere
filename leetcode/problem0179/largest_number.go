package problem0179

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 179. 最大数

// 给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数。

// 注意：输出结果可能非常大，所以你需要返回一个字符串而不是整数。

func largestNumber(nums []int) string {
	n := len(nums)
	sort.Slice(nums, func(i, j int) bool {
		numI, numJ := nums[i], nums[j]
		lenI, lenJ := len(strconv.Itoa(numI)), len(strconv.Itoa(numJ))
		iJ, jI := numI*int(math.Pow10(lenJ))+numJ, numJ*int(math.Pow10(lenI))+numI
		return iJ > jI
	})
	if nums[0] == 0 {
		return "0"
	}
	sb := &strings.Builder{}
	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(nums[i]))
	}
	return sb.String()
}
