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

	var n int
	fmt.Fscanln(reader, &n)

	for i := 0; i < n; i++ {
		oxResult := ""
		fmt.Fscanln(reader, &oxResult)

		var oCount, result int
		for _, ox := range oxResult {
			if ox == 'O' {
				oCount++
				result += oCount
			} else {
				oCount = 0
			}
		}
		fmt.Fprintln(writer, result)
	}
}
