package hj089

import (
	"fmt"
)

func Do() {
	var input = make([]string, 4)
	var scanN int
	for i := 0; i < 4; i++ {
		scanN, _ = fmt.Scan(&input[i])
	}
	for scanN > 0 {
		fmt.Println(operate(input))

		for i := 0; i < 4; i++ {
			scanN, _ = fmt.Scan(&input[i])
		}
	}
}

var weights = map[string]int{
	"2":  2,
	"3":  3,
	"4":  4,
	"5":  5,
	"6":  6,
	"7":  7,
	"8":  8,
	"9":  9,
	"10": 10,
	"J":  11,
	"Q":  12,
	"K":  13,
	"A":  1,
}
var operators = []string{"+", "-", "*", "/"}

// 采用暴力枚举方法，其中的全排列算法不是标准算法
func operate(input []string) string {
	if hasJoker(input) {
		return "ERROR"
	}
	var output = "NONE"
	var found = false
	var permutations = permute(input)
	for _, permutation := range permutations {
		if found {
			break
		}
		for i := 0; i < 4; i++ {
			if found {
				break
			}
			var pointI = compute(weights[permutation[0]], weights[permutation[1]], operators[i])
			for j := 0; j < 4; j++ {
				if found {
					break
				}
				var pointJ = compute(pointI, weights[permutation[2]], operators[j])
				for k := 0; k < 4; k++ {
					var pointK = compute(pointJ, weights[permutation[3]], operators[k])

					if pointK == 24 {
						output = permutation[0] + operators[i] +
							permutation[1] + operators[j] +
							permutation[2] + operators[k] +
							permutation[3]
						found = true
						break
					}
				}
			}
		}
	}

	return output
}

// 获取输入的全排列
func permute(input []string) [][]string {
	var result = make([][]string, 0)
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if j == i {
				continue
			}
			for k := 0; k < 4; k++ {
				if k == i || k == j {
					continue
				}
				for l := 0; l < 4; l++ {
					if l == i || l == j || l == k {
						continue
					}
					result = append(result, []string{input[i], input[j], input[k], input[l]})
				}
			}
		}
	}

	return result
}

func compute(x, y int, operator string) int {
	switch operator {
	case "+":
		return x + y
	case "-":
		return x - y
	case "*":
		return x * y
	case "/":
		return x / y
	default:
		return x + y
	}
}

func hasJoker(input []string) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == "joker" || input[i] == "JOKER" {
			return true
		}
	}
	return false
}
