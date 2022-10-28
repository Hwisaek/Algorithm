package main

import (
	b "bufio"
	. "fmt"
	"os"
)

var (
	sc = b.NewScanner(os.Stdin)
)

func main() {
	sc.Scan()

	count := 0
loop:
	for sc.Scan() {
		word := sc.Text()

		m := make(map[byte]any)
		for i := range word {
			c := word[i]
			if i == 0 {
				m[c] = nil
				continue
			}
			prev := word[i-1]
			if c != prev {
				if _, ok := m[c]; ok {
					continue loop
				}
				m[prev] = nil
			}
		}
		count++
	}
	Println(count)
}
