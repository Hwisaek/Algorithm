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

	var x, y, z int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			a := 0
			Fscan(r, &a)
			if a > x {
				x, y, z = a, i, j
			}
		}
	}

	Printf("%d\n%d %d", x, y+1, z+1)
}
