package main

import (
	b "bufio"
	. "fmt"
	"os"
)

var (
	n, k, count int
	r           = b.NewReader(os.Stdin)
)

func main() {
	Fscanln(r, &n, &k)
	s := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		Fscanln(r, &s[i])
	}

	for _, c := range s {
		if k >= c {
			count += k / c
			k %= c
			if k == 0 {
				break
			}
		}
	}
	Println(count)
}
