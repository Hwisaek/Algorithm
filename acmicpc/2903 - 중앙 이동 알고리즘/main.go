package main

import (
	. "fmt"
)

func main() {
	var N int
	Scan(&N)
	a := 1 + 1<<N
	Print(a * a)
}
