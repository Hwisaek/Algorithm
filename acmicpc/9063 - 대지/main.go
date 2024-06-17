package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N int
	Fscan(r, &N)

	x1, y1, x2, y2 := 10001, 10001, -10001, -10001
	for i := 0; i < N; i++ {
		var x, y int
		Fscan(r, &x, &y)
		if x < x1 {
			x1 = x
		}
		if y < y1 {
			y1 = y
		}
		if x > x2 {
			x2 = x
		}
		if y > y2 {
			y2 = y
		}
	}
	Fprint(w, (x2-x1)*(y2-y1))
}
