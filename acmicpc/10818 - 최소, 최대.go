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

	var n, min, max int
	fmt.Fscanln(reader, &n)

	tmp := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &tmp)
		if i == 0 {
			max = tmp
			min = tmp
			continue
		}
		if tmp > max {
			max = tmp
		}
		if tmp < min {
			min = tmp
		}
	}

	fmt.Fprintln(writer, min, max)
}
