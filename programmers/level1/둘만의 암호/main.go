package main

import (
	"fmt"
)

func solution(s string, skip string, index int) string {
	result := make([]rune, 0)

	alphabets := make([]rune, 0)
	for i := 'a'; i <= 'z'; i++ {
		alphabets = append(alphabets, i)
	}
	alphabets = append(alphabets, append(alphabets, alphabets...)...)

	runes := make(map[rune]any)
	for _, v := range skip {
		runes[v] = nil
	}

	for _, i := range s {
		c := i + rune(index)

		for _, k := range runes {
			if c > 'z' {
				c -= 'z' - 'a' + 1
			}
			if i < k && c >= k {
				c++
			}
		}

		if c > 'z' {
			c -= 'z' - 'a' + 1
		}

		result = append(result, c)
	}

	return string(result)
}

func main() {
	fmt.Println(solution("aukks", "wbqd", 5))
}
