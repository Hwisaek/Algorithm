package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

var r = bufio.NewReader(os.Stdin)
var w = bufio.NewWriter(os.Stdout)

func main() {
	defer w.Flush()

	var n int
	Fscanln(r, &n)

	a := make(map[string]struct{})
	for i := 0; i < n; i++ {
		var n, t string
		Fscanln(r, &n, &t)
		if t == "enter" {
			a[n] = struct{}{}
		} else {
			delete(a, n)
		}
	}

	s := make([]string, 0, len(a))
	for k := range a {
		s = append(s, k)
	}
	sort.Strings(s)
	for _, e := range s {
		Fprintln(w, e)
	}
}
