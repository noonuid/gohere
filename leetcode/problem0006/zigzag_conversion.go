package problem0006

import "bytes"

// 6. Z 字形变换

// 将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。

// 比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：

// P   A   H   N
// A P L S I I G
// Y   I   R
// 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。

// 请你实现这个将字符串进行指定行数变换的函数：

// string convert(string s, int numRows);

// 先枚举每一行，再枚举行中的每一个周期，直接推算出周期中的两个元素（某些行的某些周期只有一个元素）在原字符串中的索引。
func convert(s string, numRows int) string {
	n := len(s)
	if numRows <= 1 || numRows >= n {
		return s
	}

	answer := make([]byte, 0, n)
	t := numRows + numRows - 2
	for row := 0; row < numRows; row++ {
		for cycleStart := 0; cycleStart+row < n; cycleStart = cycleStart + t {
			answer = append(answer, s[cycleStart+row])
			if 0 < row && row < numRows-1 && cycleStart+t-row < n {
				answer = append(answer, s[cycleStart+t-row])
			}
		}
	}

	return string(answer)
}

// 模拟。优化空间复杂度。
func convert_simulation_optimized_space(s string, numRows int) string {
	n := len(s)
	if numRows <= 1 || numRows >= n {
		return s
	}

	matrix := make([][]byte, numRows)
	t := numRows + numRows - 2
	for row, index := 0, 0; index < n; index++ {
		matrix[row] = append(matrix[row], s[index])
		if index%t < numRows-1 {
			row++
		} else {
			row--
		}
	}

	return string(bytes.Join(matrix, nil))
}

// 模拟。
func convert_simulation(s string, numRows int) string {
	n := len(s)
	if numRows <= 1 || numRows >= n {
		return s
	}

	t := numRows + numRows - 2
	numCols := n / t * (numRows - 1)
	if remain := n % t; remain > 0 {
		if remain <= numRows {
			numCols = numCols + 1
		} else {
			numCols = numCols + remain - (numRows - 1)
		}
	}

	matrix := make([][]byte, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]byte, numCols)
	}
	for row, col, index := 0, 0, 0; index < n; index++ {
		matrix[row][col] = s[index]
		if index%t < numRows-1 {
			row++
		} else {
			row--
			col++
		}
	}

	answer := make([]byte, 0, n)
	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			if matrix[row][col] != 0 {
				answer = append(answer, matrix[row][col])
			}
		}
	}

	return string(answer)
}
