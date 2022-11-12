package main

import (
	"regexp"
)

func solution(babbling []string) (result int) {
	regex := regexp.MustCompile("aya|ye|woo|ma")
	for _, word := range babbling {
		if len(regex.ReplaceAllString(word, "")) == 0 {
			result++
		}
	}
	return
}
