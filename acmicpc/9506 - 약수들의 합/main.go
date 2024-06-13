package main

import (
	. "fmt"
)

func main() {
	for {
		n, sum := 0, 1
		Scan(&n)
		if n == -1 {
			break
		}

		s := Sprintf("%d = 1 ", n)
		for i := 2; i < n; i++ {
			if n%i == 0 {
				s += Sprintf("+ %d ", i)
				sum += i
			}
		}

		if sum == n {
			Print(s)
		} else {
			Printf("%d is NOT perfect.\n", n)
		}
	}
}
