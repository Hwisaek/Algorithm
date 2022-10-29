package main

import . "fmt"

var dial = []int{3, 3, 3, 4, 4, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 8, 9, 9, 9, 10, 10, 10, 10}

func main() {
	var s string
	_, _ = Scan(&s)

	time := 0
	for _, c := range s {
		time += dial[c-'A']
	}
	Print(time)
}
