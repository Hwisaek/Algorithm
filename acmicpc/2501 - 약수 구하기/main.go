package main

import (
	. "fmt"
)

var N, K int

func main() {
	Scan(&N, &K)
	for i := 1; i <= N; i++ {
		if N%i == 0 {
			K--
		}
		if K == 0 {
			Print(i)
			return
		}
	}
	Print(0)
}
