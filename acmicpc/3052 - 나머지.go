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

	var countMod [42]int
	var result int

	for i := 0; i < 10; i++ {
		var n int
		fmt.Fscan(reader, &n)

		countMod[n%42]++
	}

	for i := 0; i < 42; i++ {
		if countMod[i] != 0 {
			result++
		}
	}

	fmt.Fprintln(writer, result)
}
