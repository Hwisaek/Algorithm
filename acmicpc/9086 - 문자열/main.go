package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter((os.Stdout))
	defer w.Flush()

	T := 0
	Fscan(r, &T)
	for i := 0; i < T; i++ {
		a := ""
		Fscan(r, &a)
		Fprintf(w, "%c%c\n", a[0], a[len(a)-1])
	}
}
