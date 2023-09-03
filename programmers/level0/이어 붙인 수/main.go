package main

import "strconv"

func solution(num_list []int) int {
	strOdd := ""
	strEven := ""

	for _, v := range num_list {
		if v%2 == 0 {
			strEven += strconv.Itoa(v)
		} else {
			strOdd += strconv.Itoa(v)
		}
	}

	odd, _ := strconv.Atoi(strOdd)
	even, _ := strconv.Atoi(strEven)

	return odd + even
}
