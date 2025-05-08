package solution

import (
	"strings"
)

func Solution(S string) (answer int) {
	answer = -1

	passwords := strings.Split(S, " ")
	for _, password := range passwords {
		if !isValid(password) {
			continue
		}
		answer = max(answer, len(password))
	}

	return
}

func isValid(password string) bool {
	number := 0
	letter := 0
	for _, c := range password {
		switch {
		case c >= '0' && c <= '9':
			number++
		case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z':
			letter++
		default:
			return false
		}
	}

	return number%2 == 1 && letter%2 == 0
}
