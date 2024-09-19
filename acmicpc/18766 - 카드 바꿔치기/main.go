package main

import (
	"bufio"
	"fmt"
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

	T, _ := strconv.Atoi(scan())
	for i := 0; i < T; i++ {
		a := make(map[string]int)
		b := make(map[string]int)

		N, _ := strconv.Atoi(scan())
		for i := 0; i < N; i++ {
			card := scan()
			a[card]++
		}

		for i := 0; i < N; i++ {
			card := scan()
			b[card]++
		}

		cheat := false
		for k, v := range b {
			if v != a[k] {
				fmt.Fprintln(w, "CHEATER")
				cheat = true
				break
			}
		}
		if !cheat {
			fmt.Fprintln(w, "NOT CHEATER")
		}
	}
}
