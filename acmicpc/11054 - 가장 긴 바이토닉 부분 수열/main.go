package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	N := 0
	fmt.Fscanln(reader, &N)
	A := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &A[i])
	}

	d := map[int]int{}
	for _, item := range A {
		length := 0
		for j := 1; j <= len(d); j++ {
			if d[j] < item {
				length = j
			}
		}

		d[length+1] = item
	}

	fmt.Fprintln(writer, len(d))
}
