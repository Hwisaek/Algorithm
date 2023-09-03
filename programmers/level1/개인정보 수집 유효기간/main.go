package main

import (
	"strconv"
	"strings"
)

func solution(today string, terms []string, privacies []string) (result []int) {
	date := strings.Split(today, ".")
	year, _ := strconv.Atoi(date[0])
	month, _ := strconv.Atoi(date[1])
	day, _ := strconv.Atoi(date[2])
	now := year*28*12 + month*28 + day

	t := map[string]int{}

	for _, term := range terms {
		split := strings.Split(term, " ")
		i, _ := strconv.Atoi(split[1])
		t[split[0]] = i * 28
	}

	for i, privacy := range privacies {
		split := strings.Split(privacy, " ")

		date := strings.Split(split[0], ".")
		year, _ := strconv.Atoi(date[0])
		month, _ := strconv.Atoi(date[1])
		day, _ := strconv.Atoi(date[2])

		agreeDate := year*28*12 + month*28 + day
		expiredAt := agreeDate + t[split[1]]

		if expiredAt <= now {
			result = append(result, i+1)
		}

	}

	return
}

func main() {
	solution("2022.05.19", []string{"A 6", "B 12", "C 3"}, []string{"2021.05.02 A", "2021.07.01 B", "2022.02.19 C", "2022.02.20 C"})
	solution("2020.01.01", []string{"Z 3", "D 5"}, []string{"2019.01.01 D", "2019.11.15 Z", "2019.08.02 D", "2019.07.01 D", "2018.12.28 Z"})
}
