package main

import (
	"fmt"
	"strconv"
)

var answer = map[int]interface{}{}

func solution(numbers string) int {
	dfs("", numbers)
	return len(answer)
}

func isPrime(num int) bool {
	if num < 2 {
		return false
	}

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func dfs(p, s string) {
	atoi, _ := strconv.Atoi(p)
	if isPrime(atoi) {
		answer[atoi] = nil
	}

	if len(s) == 0 {
		return
	}

	for i := range s {
		dfs(p+fmt.Sprintf("%c", s[i]), fmt.Sprintf("%v%v", s[:i], s[i+1:]))
	}
}

func main() {
	solution("17")
	solution("011")
}
