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

	var a [5]string
	for i := 0; i < 5; i++ {
		Fscan(r, &a[i])
	}

	for i := 0; i < 15; i++ {
		for j := 0; j < 5; j++ {
			if len(a[j]) > i {
				Fprint(w, string(a[j][i]))
			}
		}
	}
}
