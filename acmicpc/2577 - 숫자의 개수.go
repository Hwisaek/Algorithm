package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var count [10]int

	sum := 1
	for i := 0; i < 3; i++ {
		var n int
		fmt.Fscan(reader, &n)
		sum *= n
	}

	str := strconv.Itoa(sum)
	for i := 0; i < len(str); i++ {
		num, _ := strconv.ParseInt(str[i:i+1], 10, 0)
		count[num]++
	}

	for i := 0; i < 10; i++ {
		fmt.Fprintln(writer, count[i])
	}
}
