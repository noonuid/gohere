package problem0378

// 378. 有序矩阵中第 K 小的元素

// 给你一个 n x n 矩阵 matrix ，其中每行和每列元素均按升序排序，找到矩阵中第 k 小的元素。
// 请注意，它是 排序后 的第 k 小元素，而不是第 k 个 不同 的元素。

// 你必须找到一个内存复杂度优于 O(n2) 的解决方案。

// 二分查找。
func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	getOrder := func(mid int) int {
		i, j := n-1, 0
		order := 0
		for i > -1 && j < n {
			if matrix[i][j] <= mid {
				order += i + 1
				j++
			} else {
				i--
			}
		}
		return order
	}
	left, right := matrix[0][0], matrix[n-1][n-1]
	for left < right {
		mid := left + (right-left)/2
		order := getOrder(mid)
		if order >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
