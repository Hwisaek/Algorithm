package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	N, _ := strconv.Atoi(s.Text())

	m := map[string]struct{}{"ChongChong": {}}
	for i := 0; i < N; i++ {
		s.Scan()
		a := s.Text()
		s.Scan()
		b := s.Text()

		if _, ok := m[a]; ok {
			m[b] = struct{}{}
		} else if _, ok := m[b]; ok {
			m[a] = struct{}{}
		}
	}

	fmt.Fprint(w, len(m))
}
