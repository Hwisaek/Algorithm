package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	Fscan(r, &n)

	for i := 0; i < n; i++ {
		Fprint(w, strings.Repeat(" ", n-1-i), strings.Repeat("*", i*2+1), "\n")
	}
	for i := n - 2; i >= 0; i-- {
		Fprint(w, strings.Repeat(" ", n-1-i), strings.Repeat("*", i*2+1), "\n")
	}
}
