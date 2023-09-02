package main

import "strings"

func solution(my_string string, indices []int) string {
	b := []byte(my_string)
	for _, index := range indices {
		b[index] = byte(' ')
	}
	return strings.ReplaceAll(string(b), " ", "")
}

func main() {
	solution("apporoograpemmemprs", []int{1, 16, 6, 15, 0, 10, 11, 3})
}
