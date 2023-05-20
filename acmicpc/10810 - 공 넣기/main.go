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

	var N, M, i, j, k int
	Fscan(r, &N, &M)

	m := make([]int, N+1)
	for x := 0; x < M; x++ {
		Fscan(r, &i, &j, &k)
		for y := i; y <= j; y++ {
			m[y] = k
		}
	}

	for x := 1; x <= N; x++ {
		Fprint(w, m[x], " ")
	}
}
