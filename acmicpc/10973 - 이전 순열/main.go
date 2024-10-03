package main

import (
	"bufio"
	"fmt"
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
	N, mask      int
	tgt          []string
	arr          []string
	a            []string
	last, target string
)

func main() {
	defer w.Flush()
	s.Split(bufio.ScanWords)

	N, _ = strconv.Atoi(scan())

	arr = make([]string, N)
	tgt = make([]string, N)
	a = make([]string, N)
	for i := 0; i < N; i++ {
		tgt[i] = scan()
		arr[i] = strconv.Itoa(i + 1)
	}

	target = strings.Join(tgt, " ")
	if strings.Join(arr, " ") == target {
		fmt.Fprintln(w, -1)
		return
	}

	f(N)

}
func f(depth int) {
	if depth == 0 {
		now := strings.Join(a, " ")
		if now == target {
			fmt.Fprintln(w, last)
			w.Flush()
			os.Exit(0)
		}
		last = now
		return
	}

	for i := 0; i < N; i++ {
		if mask&(1<<i) > 0 {
			continue
		}

		mask |= 1 << i
		a[N-depth] = arr[i]
		f(depth - 1)
		mask ^= 1 << i
	}
}
