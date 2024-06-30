package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
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

	var N, M int
	N, _ = strconv.Atoi(scan())
	M, _ = strconv.Atoi(scan())

	d := make(map[string]int)
	for i := 0; i < N; i++ {
		w := scan()
		if len(w) < M {
			continue
		}
		d[w]++
	}

	dl := make([][]any, 0, len(d))
	for k, v := range d {
		dl = append(dl, []any{k, v})
	}

	slices.SortFunc(dl, func(a, b []any) int {
		i1 := a[1].(int)
		j1 := b[1].(int)
		if i1 > j1 {
			return -1
		} else if i1 == j1 {
			i0 := a[0].(string)
			j0 := b[0].(string)
			if len(i0) > len(j0) {
				return -1
			} else if len(i0) < len(j0) {
				return 1
			}
			if i0 < j0 {
				return -1
			} else if i0 > j0 {
				return 1
			}
		} else {
			return 1
		}
		return 0
	})

	for _, e := range dl {
		fmt.Fprintln(w, e[0])
	}
}
