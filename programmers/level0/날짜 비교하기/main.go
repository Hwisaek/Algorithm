package main

import (
	"time"
)

func solution(date1 []int, date2 []int) int {
	time1 := time.Date(date1[0], time.Month(date1[1]), date1[2], 0, 0, 0, 0, time.Local)
	time2 := time.Date(date2[0], time.Month(date2[1]), date2[2], 0, 0, 0, 0, time.Local)
	if time1.Before(time2) {
		return 1
	}
	return 0
}

func main() {
	solution([]int{2021, 12, 8}, []int{2021, 12, 19})
}
