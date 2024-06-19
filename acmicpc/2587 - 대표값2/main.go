package main

import (
	. "fmt"
	"sort"
)

func main() {
	a := make([]int, 5)
	for i := 0; i < 5; i++ {
		Scanln(&a[i])
	}
	sort.Ints(a)
	Println((a[0] + a[1] + a[2] + a[3] + a[4]) / 5)
	Println(a[2])
}
