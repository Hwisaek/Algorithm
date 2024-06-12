package main

import (
	. "fmt"
	"strconv"
)

func main() {
	N, B := "", 0
	Scan(&N, &B)
	i, _ := strconv.ParseInt(N, B, 32)
	Print(i)
}
