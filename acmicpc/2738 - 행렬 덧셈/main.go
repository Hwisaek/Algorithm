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

	var n, m int
	Fscanln(r, &n, &m)

	C := [100][100]int{}

	for k := 0; k < 2; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				var a int
				Fscan(r, &a)
				C[i][j] += a
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			Fprint(w, C[i][j], " ")
		}
		Fprintln(w)
	}
}
