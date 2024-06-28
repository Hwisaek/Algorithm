package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(math.Pow(2, float64(N)))
}
