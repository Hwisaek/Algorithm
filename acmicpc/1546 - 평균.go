package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, max int
	var scores []int
	var avg float64

	fmt.Fscan(reader, &n)
	scores = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &scores[i])
		if scores[i] > max {
			max = scores[i]
		}
	}

	for _, score := range scores {
		avg += float64(score) / float64(max) * 100
	}
	avg /= float64(n)

	fmt.Fprintln(writer, avg)
}
