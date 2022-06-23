package 자릿수_더하기

import (
	"strconv"
)

func solution(n int) (result int) {
	itoa := strconv.Itoa(n)

	for _, c := range itoa {
		num := c - '0'
		result += int(num)
	}
	return
}
