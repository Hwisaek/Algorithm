package main

import "strconv"

func solution(age int) (result string) {
	for _, c := range strconv.Itoa(age) {
		result += string(c - '0' + 'a')
	}
	return
}

func main() {
	solution(23)
}
