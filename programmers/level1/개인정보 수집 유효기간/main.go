package main

import (
	"strconv"
	"strings"
)

func solution(today string, terms []string, privacies []string) []int {
	t := map[string]int{}
	p := map[string]string{}

	for _, term := range terms {
		split := strings.Split(term, " ")
		i, _ := strconv.Atoi(split[1])
		t[split[0]] = i
	}

	for _, privacy := range privacies {
		split := strings.Split(privacy, " ")
		p[split[0]] = split[1]
	}

	return
}
