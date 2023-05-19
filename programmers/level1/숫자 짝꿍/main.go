package main

import (
	"sort"
	"strconv"
	"strings"
)

func solution(X string, Y string) (result string) {
	x := make(map[string]int)
	for _, v := range X {
		x[string(v)]++
	}
	y := make(map[string]int)
	for _, v := range Y {
		y[string(v)]++
	}

	strings.count
	list := make([]string, 0)
	for k := range x {
		if _, ok := y[k]; ok {
			for x[k] > 0 && y[k] > 0 {
				list = append(list, k)
				x[k]--
				y[k]--
			}
		}
	}

	if len(list) == 0 {
		return "-1"
	}

	sort.Slice(list, func(i, j int) bool { return list[i] > list[j] })

	result = strings.Join(list, "")

	i, _ := strconv.Atoi(result)
	if i == 0 {
		return "0"
	}

	return
}

func main() {
	solution("100", "2345")
}
