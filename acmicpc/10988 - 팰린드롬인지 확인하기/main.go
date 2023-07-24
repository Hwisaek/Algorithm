package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n string
	Fscan(r, &n)

	split := strings.Split(n, "")
	sort.Sort(sort.Reverse(sort.StringSlice(split)))

	if n == reverse(n) {
		Fprint(w, 1)
		return
	}
	Fprint(w, 0)
}

func reverse(s string) (r string) {
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return
}
