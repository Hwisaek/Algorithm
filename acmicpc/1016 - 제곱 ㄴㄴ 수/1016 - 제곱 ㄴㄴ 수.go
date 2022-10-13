package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	var min, max int
	Fscan(r, &min, &max)

	isPow := make([]bool, max-min+11)

	count := max - min + 1

	iMax := int(math.Sqrt(float64(max)))
	for i := 2; i <= iMax; i++ {
		ii := i * i
		x := min / ii
		if x*ii < min {
			x++
		}
		for p := ii * x; p <= max; p += ii {
			if p >= min {
				if !isPow[p-min] {
					count--
				}
				isPow[p-min] = true
			}
		}
	}
	Fprintln(w, count)
	w.Flush()
}
