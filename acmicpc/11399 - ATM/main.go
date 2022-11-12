package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var (
	r = bufio.NewReader(os.Stdin)
)

func main() {
	var n, result, sum int
	Fscan(r, &n)

	p := make([]int, n)
	for i := range p {
		Fscan(r, &p[i])
	}

	sort.Ints(p)
	for _, m := range p {
		sum += m
		result += sum
	}

	Print(result)
}
