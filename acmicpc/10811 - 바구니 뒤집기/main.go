package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)
	var N, M int
	Fscan(r, &N, &M)

	b := make([]int, N)
	for i := 0; i < N; i++ {
		b[i] = i + 1
	}

	for k := 0; k < M; k++ {
		var i, j int
		Fscan(r, &i, &j)

		for x, y := i-1, j-1; x < y; x, y = x+1, y-1 {
			b[x], b[y] = b[y], b[x]
		}
	}

	for _, n := range b {
		Fprint(w, n, " ")
	}
	w.Flush()
}
