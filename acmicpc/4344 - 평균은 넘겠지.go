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

	var C int
	fmt.Fscanln(reader, &C)

	for i := 0; i < C; i++ {
		var N, count int
		var avg float64
		fmt.Fscan(reader, &N)
		scores := make([]float64, N)
		for i := 0; i < N; i++ {
			fmt.Fscan(reader, &scores[i])
			avg += scores[i]
		}
		avg /= float64(N)

		for i := 0; i < N; i++ {
			if scores[i] > avg {
				count++
			}
		}
		fmt.Fprintf(writer, "%.3f%%\n", float64(count)/float64(N)*100)
	}
}
