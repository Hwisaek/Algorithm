package main

import (
	. "fmt"
)

func main() {
	var n int
	Scan(&n)
	Printf("%d\n%d", n*(n-1)/2, 2)
}
