package main

import (
	"fmt"
	"sort"
)

func solution(s string, skip string, index int) string {
	result := make([]rune, 0)

	runes := []rune(skip)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

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
