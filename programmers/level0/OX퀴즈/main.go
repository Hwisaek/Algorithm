package main

import (
	"strconv"
	"strings"
)

func solution(quiz []string) (result []string) {
	for _, e := range quiz {
		split := strings.Split(e, " = ")
		answer, _ := strconv.Atoi(split[1])
		exp := strings.Split(split[0], " ")
		x, y := exp[0], exp[2]

		switch exp[1] {
		case "+":
			nx, _ := strconv.Atoi(x)
			ny, _ := strconv.Atoi(y)
			if nx+ny == answer {
				result = append(result, "O")
			} else {
				result = append(result, "X")
			}
		case "-":
			nx, _ := strconv.Atoi(x)
			ny, _ := strconv.Atoi(y)
			if nx-ny == answer {
				result = append(result, "O")
			} else {
				result = append(result, "X")
			}
		}
	}
	return
}
