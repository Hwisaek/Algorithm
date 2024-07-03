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

var (
	visited []bool
	arr     []string
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)
	N, _ := strconv.Atoi(scan())
	M, _ := strconv.Atoi(scan())

	visited = make([]bool, N+1)
	arr = make([]string, M)
	backtrack(N, M, []string{})
}

func backtrack(N, M int, arr []string) {
	if len(arr) >= M {
		w.WriteString(strings.Join(arr, " "))
		w.WriteString("\n")
		return
	}
	for i := 1; i <= N; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		backtrack(N, M, append(arr, strconv.Itoa(i)))
		visited[i] = false
	}
}
