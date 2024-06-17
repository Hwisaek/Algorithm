package main

import (
	. "fmt"
	"sort"
)

func main() {
	n := make([]int, 3)
	Scanln(&n[0], &n[1], &n[2])
	sort.Ints(n)
	if n[2] >= n[0]+n[1] {
		n[2] = n[0] + n[1] - 1
	}
	Println(n[0] + n[1] + n[2])
}
