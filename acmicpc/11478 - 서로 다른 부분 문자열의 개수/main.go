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

	var S string
	Fscan(r, &S)

	m := map[string]struct{}{}
	for i := 0; i < len(S); i++ {
		for j := 1; j <= len(S)-i; j++ {
			m[S[i:i+j]] = struct{}{}
		}
	}

	Fprint(w, len(m))
}
