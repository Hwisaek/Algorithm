package main

import . "fmt"

func main() {
	var X, N, a, b, i int
	Scan(&X, &N)
	for ; i < N; i++ {
		Scan(&a, &b)
		X -= a * b
	}
	r := "Yes"
	if X != 0 {
		r = "No"
	}
	Print(r)
}
