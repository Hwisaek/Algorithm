package main

import (
	b "bufio"
	. "fmt"
	"os"
	"sort"
)

var (
	r = b.NewReader(os.Stdin)
	c = -1

	n, count int
)

func main() {
	Fscan(r, &n)
	a := make([][2]int, n)
	for i := 0; i < n; i++ {
		a[i] = [2]int{0, 0}
		Fscan(r, &a[i][0], &a[i][1])
	}

	sort.Slice(a, func(i, j int) bool {
		if a[i][1] == a[j][1] {
			return a[i][0] < a[j][0]
		}
		return a[i][1] < a[j][1]
	})

	for _, ints := range a {
		if c <= ints[0] {
			c = ints[1]
			count++
		}
	}
	Print(count)
}
