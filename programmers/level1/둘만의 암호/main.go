package main

import (
	"fmt"
	"strings"
)

func solution(s string, skip string, index int) (result string) {
	aMap := make(map[rune]int)
	aSlice := make([]rune, 0, 60)

	for i := 'a'; i <= 'z'; i++ {
		if strings.IndexRune(skip, i) != -1 {
			continue
		}
		aMap[i] = len(aSlice)
		aSlice = append(aSlice, i)
	}

	aSlice = append(aSlice, append(aSlice, aSlice...)...)

	var sb strings.Builder
	for _, v := range s {
		sb.WriteRune(aSlice[aMap[v]+index])
	}

	return sb.String()
}

func main() {
	fmt.Println(solution("aukks", "wbqd", 5))
}
