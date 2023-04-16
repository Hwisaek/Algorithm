package main

import (
	"strconv"
	"strings"
)

func solution(babbling []string) (result int) {
	words := []string{"aya", "ye", "woo", "ma"}

babbling:
	for _, origin := range babbling {
		w := origin
		for i, word := range words {
			w = strings.ReplaceAll(w, word, strconv.Itoa(i))
		}

		if _, err := strconv.Atoi(w); err == nil {
			for i := 1; i < len(w); i++ {
				if w[i] == w[i-1] {
					continue babbling
				}
			}
			result++
		}
	}
	return
}

func main() {
	// solution([]string{"aya", "yee", "u", "maa"})
	solution([]string{"ayaye", "uuu", "yeye", "yemawoo", "ayaayaa"})
}
