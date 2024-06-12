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

	a := [100][100]bool{}

	var n, s int
	Fscan(r, &n)

	for i := 0; i < n; i++ {
		x, y := 0, 0
		Fscan(r, &x, &y)
		for j := x; j < x+10; j++ {
			for k := y; k < y+10; k++ {
				if !a[j][k] {
					a[j][k] = true
					s++
				}
			}
		}
	}

	Fprint(w, s)
}
