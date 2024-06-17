package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for {
		a := make([]int, 3)
		Fscan(r, &a[0], &a[1], &a[2])

		if a[0] == a[1] && a[1] == a[2] && a[2] == 0 {
			break
		}

		sort.Ints(a)

		if a[0] == a[1] && a[1] == a[2] {
			Fprintln(w, "Equilateral")
		} else if a[0]+a[1] <= a[2] {
			Fprintln(w, "Invalid")
		} else if a[0] == a[1] || a[1] == a[2] {
			Fprintln(w, "Isosceles")
		} else {
			Fprintln(w, "Scalene")
		}
	}
}
