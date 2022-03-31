package hj094

import "fmt"

func Do() {
	var canditateNumber, voterNumber, scanN int
	fmt.Scan(&canditateNumber)
	var candidates = make([]string, canditateNumber)
	for i := 0; i < canditateNumber; i++ {
		fmt.Scan(&candidates[i])
	}
	fmt.Scan(&voterNumber)
	var tickets = make([]string, voterNumber)
	for i := 0; i < voterNumber; i++ {
		scanN, _ = fmt.Scan(&tickets[i])
	}

	for scanN > 0 {
		var scores = make(map[string]int, canditateNumber)
		var invalidTicket = 0
		for _, ticket := range tickets {
			if isValidTicket(candidates, ticket) {
				scores[ticket]++
			} else {
				invalidTicket++
			}
		}

		for _, candidate := range candidates {
			fmt.Printf("%s : %d\n", candidate, scores[candidate])
		}
		fmt.Printf("Invalid : %d\n", invalidTicket)

		fmt.Scan(&canditateNumber)
		candidates = make([]string, canditateNumber)
		for i := 0; i < canditateNumber; i++ {
			fmt.Scan(&candidates[i])
		}
		fmt.Scan(&voterNumber)
		tickets = make([]string, voterNumber)
		for i := 0; i < voterNumber; i++ {
			scanN, _ = fmt.Scan(&tickets[i])
		}
	}
}

func isValidTicket(candidates []string, ticket string) bool {
	for _, candidate := range candidates {
		if candidate == ticket {
			return true
		}
	}
	return false
}
