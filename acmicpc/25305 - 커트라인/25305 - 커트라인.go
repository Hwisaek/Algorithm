package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)

func main() {
	var N, k int
	Scan(&N, &k)

	x := make([]int, N)
	for i := 0; i < N; i++ {
		Fscan(r, &x[i])
	}
	sort.Ints(x)
	Println(x[len(x)-k])
}
