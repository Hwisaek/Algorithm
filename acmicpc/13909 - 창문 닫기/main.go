package main

import (
	. "fmt"
	"math"
)

func main() {
	var n int
	Scan(&n)
	Print(int(math.Sqrt(float64(n))))
}
