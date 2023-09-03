package main

import "strings"

func solution(my_string string) (result []string) {
	for _, v := range strings.Split(my_string, " ") {
		if v != "" {
			result = append(result, v)
		}
	}
	return
}

func main() {
	solution(" i    love  you")
	solution("    programmers  ")
}
