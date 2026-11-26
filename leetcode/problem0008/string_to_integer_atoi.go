package problem0008

import "math"

func myAtoi(s string) int {
	n, i := len(s), 0
	for i < n && s[i] == ' ' {
		i++
	}
	sign := 1
	if i < n && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1
		}
		i++
	}
	abs := 0
	for i < n && '0' <= s[i] && s[i] <= '9' {
		digit := int(s[i] - '0')
		if abs > (math.MaxInt32-digit)/10 {
			if sign == 1 {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		abs = abs*10 + digit
		i++
	}
	return sign * abs
}

type automaton struct {
	sign  int
	abs   int
	state string
	table map[string][]string
}

func newAutomaton() *automaton {
	return &automaton{
		sign:  1,
		abs:   0,
		state: "start",
		table: map[string][]string{
			"start":     {"start", "signed", "in_number", "end"},
			"signed":    {"end", "end", "in_number", "end"},
			"in_number": {"end", "end", "in_number", "end"},
			"end":       {"end", "end", "end", "end"},
		},
	}
}

func (a *automaton) run(ch byte) {
	a.state = a.table[a.state][a.get_col(ch)]
	switch a.state {
	case "signed":
		if ch == '-' {
			a.sign = -1
		}
	case "in_number":
		digit := int(ch - '0')
		if a.abs > (math.MaxInt32-digit)/10 {
			if a.sign == 1 {
				a.abs = math.MaxInt32
			} else {
				a.abs = math.MaxInt32 + 1
			}
			a.state = "end"
		} else {
			a.abs = a.abs*10 + digit
		}
	}
}

func (a *automaton) get_col(ch byte) int {
	switch ch {
	case ' ':
		return 0
	case '-', '+':
		return 1
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return 2
	default:
		return 3
	}
}

func myAtoi_automaton(s string) int {
	auto := newAutomaton()
	n := len(s)
	for i := 0; i < n; i++ {
		auto.run(s[i])
	}
	return auto.sign * auto.abs
}
