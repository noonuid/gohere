package hj029

import "fmt"

func Do() {
	var input1, input2 string
	var n1, _ = fmt.Scan(&input1)
	var n2, _ = fmt.Scan(&input2)
	for n1 > 0 && n2 > 0 {
		fmt.Println(encrypt(input1))
		fmt.Println(decrypt(input2))

		n1, _ = fmt.Scan(&input1)
		n2, _ = fmt.Scan(&input2)
	}
}

// 字符串加密
func encrypt(input string) string {
	var bytes = []byte(input)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] = 'a' + (bytes[i]-'A'+1)%26
		} else if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = 'A' + (bytes[i]-'a'+1)%26
		} else if bytes[i] >= '0' && bytes[i] <= '9' {
			bytes[i] = '0' + (bytes[i]-'0'+1)%10
		}
	}

	return string(bytes)
}

// 字符串解密
func decrypt(input string) string {
	var bytes = []byte(input)

	for i := 0; i < len(bytes); i++ {
		if bytes[i] >= 'A' && bytes[i] <= 'Z' {
			bytes[i] = 'a' + (bytes[i]-'A'+25)%26
		} else if bytes[i] >= 'a' && bytes[i] <= 'z' {
			bytes[i] = 'A' + (bytes[i]-'a'+25)%26
		} else if bytes[i] >= '0' && bytes[i] <= '9' {
			bytes[i] = '0' + (bytes[i]-'0'+9)%10
		}
	}

	return string(bytes)
}
