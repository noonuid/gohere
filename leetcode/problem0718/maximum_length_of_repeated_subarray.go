package problem0718

// 解法1：动态规划
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	f := make([][]int, m+1)
	for i := 0; i < len(f); i++ {
		f[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if nums1[i-1] == nums2[j-1] {
				f[i][j] = f[i-1][j-1] + 1
				if result < f[i][j] {
					result = f[i][j]
				}
			} else {
				f[i][j] = 0
			}
		}
	}

	return result
}

// 解法2：滑动窗口
func findLengthSlidingWindow(nums1 []int, nums2 []int) int {
	// 分别代表长短数组
	var numsLong, numsShort []int
	if len(nums1) > len(nums2) {
		numsLong = nums1
		numsShort = nums2
	} else {
		numsLong = nums2
		numsShort = nums1
	}

	maxLength := 0
	for i := 0; i < len(numsShort)-1; i++ {
		index1, index2 := 0, len(numsShort)-1-i
		maxLength = max(maxLength, maxLengthInWindow(numsLong, numsShort, index1, index2, i+1))
	}
	for i := 0; i < len(numsLong)-len(numsShort)+1; i++ {
		index1, index2 := i, 0
		maxLength = max(maxLength, maxLengthInWindow(numsLong, numsShort, index1, index2, len(numsShort)))
	}
	for i := 0; i < len(numsShort)-1; i++ {
		index1, index2 := len(numsLong)-len(numsShort)+1+i, 0
		maxLength = max(maxLength, maxLengthInWindow(numsLong, numsShort, index1, index2, len(numsShort)-1-i))
	}

	return maxLength
}

func maxLengthInWindow(numsLong, numsShort []int, index1, index2, length int) int {
	maxLength, curLength := 0, 0
	for i := 0; i < length; i++ {
		if numsLong[index1+i] == numsShort[index2+i] {
			curLength++
		} else {
			curLength = 0
		}

		if maxLength < curLength {
			maxLength = curLength
		}
	}

	return maxLength
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}
