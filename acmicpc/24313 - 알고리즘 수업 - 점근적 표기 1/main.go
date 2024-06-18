package main

import (
	. "fmt"
)

func main() {
	var a1, a0, c, n0 int
	Scanf("%d %d\n%d\n%d", &a1, &a0, &c, &n0)

	if a1*n0+a0 <= c*n0 && a1*100+a0 <= c*100 {
		Print(1)
	} else {
		Print(0)
	}
}
