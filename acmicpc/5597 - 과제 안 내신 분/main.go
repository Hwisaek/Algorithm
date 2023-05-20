package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r, w := bufio.NewReader(os.Stdin), bufio.NewWriter(os.Stdout)

	m := make([]int, 30)
	for i, n := 0, 0; i < 28; i++ {
		Fscan(r, &n)
		m[n-1]++
	}
	for i, v := range m {
		if v == 0 {
			Fprintln(w, i+1)
		}
	}
	w.Flush()
}
