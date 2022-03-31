package nc029

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 * @param target int整型
 * @param array int整型二维数组
 * @return bool布尔型
 */
func Find(target int, array [][]int) bool {
	var m, n = len(array), len(array[0])
	var x, y = 0, n - 1
	for x <= m-1 && y >= 0 {
		if array[x][y] == target {
			return true
		} else if array[x][y] > target {
			y--
		} else if array[x][y] < target {
			x++
		}
	}
	return false
}
