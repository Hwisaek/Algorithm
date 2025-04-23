package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()

	var K, N int
	fmt.Fscanln(reader, &K, &N)

	var low = 1
	var high = 0

	s := make([]int, K)

	for i := 0; i < K; i++ {
		fmt.Fscanln(reader, &s[i])
		if s[i] > high {
			high = s[i]
		}
	}

	answer := 0
	for low <= high {
		mid := (low + high) / 2

		total := 0

		for _, e := range s {
			total += e / mid
		}

		if total >= N {
			answer = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	fmt.Fprintln(writer, answer)
}
