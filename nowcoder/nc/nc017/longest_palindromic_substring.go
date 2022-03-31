package nc017

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 *
 *
 * @param A string字符串
 * @return int整型
 */
func getLongestPalindrome(A string) int {
	// write code here
	var length = len(A)
	if length <= 1 {
		return length
	}

	var f = make([][]bool, length)
	for i := 0; i < length; i++ {
		f[i] = make([]bool, length)
		f[i][i] = true
	}

	var begin, end = 0, 0
	for subLength := 2; subLength <= length; subLength++ {
		for beginIndex := 0; beginIndex < length; beginIndex++ {
			var endIndex = beginIndex + subLength - 1

			if endIndex >= length {
				break
			}

			if A[beginIndex] == A[endIndex] {
				if subLength <= 3 {
					f[beginIndex][endIndex] = true
				} else {
					f[beginIndex][endIndex] = f[beginIndex+1][endIndex-1]
				}

				if f[beginIndex][endIndex] && subLength > end-begin+1 {
					begin = beginIndex
					end = endIndex
				}
			} else {
				f[beginIndex][endIndex] = false
			}
		}
	}

	return end - begin + 1
}
