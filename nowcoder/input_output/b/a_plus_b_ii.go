package b

import (
	"fmt"
)

func main() {
	var t, a, b int
	fmt.Scan(&t)
	for t > 0 {
		fmt.Scan(&a, &b)
		fmt.Println(a + b)

		t--
	}
}
