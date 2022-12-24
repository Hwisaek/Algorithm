package main

import (
	"strconv"
	"strings"
)

func solution(n int) (count int) {
	answer := 0
	for answer != n {
		count++
		if count%3 == 0 {
			continue
		}
		if strings.Contains(strconv.Itoa(count), "3") {
			continue
		}
		answer++
	}
	return
}

func main() {
	solution(15)
	solution(40)
}
