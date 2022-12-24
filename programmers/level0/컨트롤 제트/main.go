package main

import (
	"strconv"
	"strings"
)

func solution(s string) (result int) {
	var arr []string
	for _, i := range strings.Split(s, " ") {
		switch i {
		case "Z":
			arr = arr[:len(arr)-1]
		default:
			arr = append(arr, i)
		}
	}
	for _, e := range arr {
		n, _ := strconv.Atoi(e)
		result += n
	}
	return
}

func main() {
	solution("1 2 Z 3")
	solution("10 20 30 40")
	solution("10 Z 20 Z 1")
	solution("10 Z 20 Z")
	solution("-1 -2 -3 Z")
}
