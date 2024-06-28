package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	N := 0
	fmt.Fscan(r, &N)

	a := 0
	var m map[string]struct{}
	for name, i := "", 0; i < N; i++ {
		fmt.Fscan(r, &name)
		if name == "ENTER" {
			m = make(map[string]struct{})
			continue
		}

		if _, ok := m[name]; ok {
			continue
		}

		m[name] = struct{}{}
		a++
	}

	fmt.Fprint(w, a)
}
