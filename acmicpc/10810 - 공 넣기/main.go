package main

import (
	. "fmt"
)

func main() {
	var N, x, v int
	m := map[int]int{}
	Scan(&N)
	for i := 0; i < N; i++ {
		Scan(&x)
		m[x]++
	}
	Scan(&v)
	m[v] += 0
	Print(m[v])
}
