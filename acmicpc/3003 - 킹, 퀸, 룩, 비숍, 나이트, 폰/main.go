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

	var a, b, c, d, e, f int
	Fscan(r, &a, &b, &c, &d, &e, &f)

	result := []int{1 - a, 1 - b, 2 - c, 2 - d, 2 - e, 8 - f}

	for _, i := range result {
		Fprint(w, strconv.Itoa(i)+" ")
	}
}
