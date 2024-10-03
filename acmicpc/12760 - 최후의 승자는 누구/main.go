package main

import (
	"bufio"
	"os"
	"strconv"
)

var (
	s = bufio.NewScanner(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func scan() string {
	s.Scan()
	return s.Text()
}

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ := strconv.Atoi(scan())
	M, _ := strconv.Atoi(scan())

	arr := make([][]int, 0, N+1)
	for i := 0; i < N; i++ {
		arr[i] = make([]int, 0, M+1)
		for j := 0; j < M; j++ {
			n, _ := strconv.Atoi(scan())
			arr[i][j] = n
		}
	}
}
