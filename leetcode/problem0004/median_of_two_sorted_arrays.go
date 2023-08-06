package problem0004

// 4. 寻找两个正序数组的中位数

// 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

// 算法的时间复杂度应该为 O(log (m+n)) 。

// 遍历一半。
func findMedianSortedArrays_half_traversal(nums1 []int, nums2 []int) float64 {
	len1, len2, index1, index2, index := len(nums1), len(nums2), 0, 0, 0
	length, halfLen := len1+len2, (len1+len2)/2+1
	pre, cur := 0, 0
	for index < halfLen {
		pre = cur
		if index1 < len1 && (index2 >= len2 || nums1[index1] < nums2[index2]) {
			cur = nums1[index1]
			index, index1 = index+1, index1+1
		} else {
			cur = nums2[index2]
			index, index2 = index+1, index2+1
		}
	}
	if length%2 == 1 {
		return float64(cur)
	} else {
		return (float64(pre) + float64(cur)) / 2
	}
}

// 合并为一个数组。
func findMedianSortedArrays_merging(nums1 []int, nums2 []int) float64 {
	len1, len2, index1, index2, index := len(nums1), len(nums2), 0, 0, 0
	length := len1 + len2
	merged := make([]int, length)
	for index < length {
		if index1 < len1 && (index2 >= len2 || nums1[index1] < nums2[index2]) {
			merged[index] = nums1[index1]
			index, index1 = index+1, index1+1
		} else {
			merged[index] = nums2[index2]
			index, index2 = index+1, index2+1
		}
	}
	if length%2 == 1 {
		return float64(merged[length/2])
	} else {
		return (float64(merged[length/2-1]) + float64(merged[length/2])) / 2
	}
}
