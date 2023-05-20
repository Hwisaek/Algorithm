package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	var N, M, i, j int
	Fscan(r, &N, &M)

	l := make([]int, N+1)
	for i := 1; i <= N; i++ {
		l[i] = i
	}

	for a := 0; a < M; a++ {
		Fscan(r, &i, &j)
		l[i], l[j] = l[j], l[i]
	}

	for i := 1; i <= N; i++ {
		Fprint(w, l[i], " ")
	}
	w.Flush()
}
