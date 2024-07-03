package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
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

	backtrack(1, N, M, []string{})
}

func backtrack(i, N, M int, arr []string) {
	if len(arr) >= M {
		w.WriteString(strings.Join(arr, " "))
		w.WriteString("\n")
		return
	}
	for ; i <= N; i++ {
		backtrack(i+1, N, M, append(arr, strconv.Itoa(i)))
	}
}
