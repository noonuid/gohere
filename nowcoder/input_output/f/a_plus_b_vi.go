package f

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input string
	reader := bufio.NewReader(os.Stdin)
	for {
		bytes, _, _ := reader.ReadLine()
		if len(bytes) == 0 {
			break
		}
		input = string(bytes)
		nums := strings.Split(input, " ")
		sum := 0
		for i := 1; i < len(nums); i++ {
			num, _ := strconv.Atoi(nums[i])
			sum += num
		}
		fmt.Println(sum)
	}
}
