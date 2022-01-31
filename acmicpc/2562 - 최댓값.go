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

	var idx, max int

	for i := 0; i < 9; i++ {
		var n int
		fmt.Fscan(reader, &n)

		if n > max {
			max = n
			idx = i + 1
		}

	}

	fmt.Fprintln(writer, max)
	fmt.Fprintln(writer, idx)
}
