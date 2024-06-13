package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var T int
	Fscan(r, &T)

	for i := 0; i < T; i++ {
		var C int
		Fscan(r, &C)
		Fprint(w, C/25, " ", C%25/10, " ", C%25%10/5, " ", C%25%10%5, "\n")
	}
}
