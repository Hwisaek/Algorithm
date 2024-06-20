package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

var s, w = bufio.NewScanner(os.Stdin), bufio.NewWriter(os.Stdout)

func main() {
	s.Split(bufio.ScanWords)

	var n int
	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	hash := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		s.Scan()
		name := s.Text()
		s.Scan()
		status := s.Text()

		if status[0] == 'e' {
			hash[name] = struct{}{}
		} else {
			delete(hash, name)
		}
	}

	arr := make([]string, 0, len(hash))
	for name := range hash {
		arr = append(arr, name)
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	for _, name := range arr {
		Fprintln(w, name)
	}
	w.Flush()
}
