package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

var (
	N, M, mask int
	arr        []int
	result     []any
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())
	M, _ = strconv.Atoi(scan())

	arr = make([]int, N)
	result = make([]any, M)
	for i := 0; i < N; i++ {
		n, _ := strconv.Atoi(scan())
		arr[i] = n
	}

	sort.Ints(arr)

	f(0)
}

func f(depth int) {
	if depth == M {
		fmt.Fprintln(w, result...)
		return
	}

	for i := 0; i < N; i++ {
		if mask&(1<<i) != 0 {
			continue
		}

		mask |= 1 << i

		result[depth] = arr[i]
		f(depth + 1)

		mask ^= 1 << i
	}
}
