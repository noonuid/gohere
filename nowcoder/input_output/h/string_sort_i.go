package h

import (
	"fmt"
	"sort"
)

func main() {
	count := 0
	fmt.Scan(&count)
	strs := make([]string, count)
	for i := 0; i < count; i++ {
		n, _ := fmt.Scanf("%s", &strs[i])
		if n == 0 {
			break
		}
	}
	sort.Strings(strs)
	for i := 0; i < count; i++ {
		fmt.Printf("%s ", strs[i])
	}
}
