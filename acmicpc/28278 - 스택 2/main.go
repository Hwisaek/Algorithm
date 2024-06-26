package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	Fscan(r, &n)

	x := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var c string
		Fscan(r, &c)

		switch c {
		case "1":
			b := 0
			Fscan(r, &b)
			x = append(x, b)
		case "2":
			if len(x) > 0 {
				w.WriteString(strconv.Itoa(x[len(x)-1]) + "\n")
				x = x[:len(x)-1]
			} else {
				w.WriteString("-1\n")
			}
		case "3":
			Fprintln(w, len(x))
		case "4":
			if len(x) > 0 {
				w.WriteString("0\n")
			} else {
				w.WriteString("1\n")
			}
		case "5":
			if len(x) > 0 {
				Fprintln(w, x[len(x)-1])
			} else {
				w.WriteString("-1\n")
			}
		}
	}
}
