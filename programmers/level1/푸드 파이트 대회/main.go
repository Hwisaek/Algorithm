package main

import (
	"strconv"
	"strings"
)

func solution(food []int) (result string) {
	left := ""
	for i := 1; i < len(food); i++ {
		left += strings.Repeat(strconv.Itoa(i), food[i]/2)
	}

	var rightSb strings.Builder
	for i := len(left) - 1; i >= 0; i-- {
		rightSb.WriteByte(left[i])
	}

	return left + "0" + rightSb.String()
}
