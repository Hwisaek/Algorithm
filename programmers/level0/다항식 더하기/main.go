package main

import (
	"fmt"
	"strconv"
	"strings"
)

func solution(polynomial string) string {
	c := [2]int{}
	for _, e := range strings.Split(polynomial, " + ") {
		n, err := strconv.Atoi(e)
		if err != nil {
			n, err := strconv.Atoi(e[:len(e)-1])
			if err != nil {
				c[0]++
			} else {
				c[0] += n
			}
		}
		c[1] += n
	}
	var x []string
	if c[0] != 0 {
		if c[0] == 1 {
			x = append(x, fmt.Sprintf("x"))
		} else {
			x = append(x, fmt.Sprintf("%dx", c[0]))
		}
	}
	if c[1] != 0 {
		x = append(x, fmt.Sprintf("%d", c[1]))
	}
	return strings.Join(x, " + ")
}
