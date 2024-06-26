package main

import (
	"bufio"
	"container/ring"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	r := ring.New(n)
	s.Scan()
	for i, e := range strings.Split(s.Text(), " ") {
		y, _ := strconv.Atoi(e)
		r.Value = []int{i + 1, y}
		r = r.Next()
	}

	for ; n > 0; n-- {
		e := r.Value.([]int)
		i, y := e[0], e[1]
		fmt.Fprint(w, i, " ")
		r = r.Prev()
		r.Unlink(1)
		if y < 0 {
			y++
		}
		r = r.Move(y)
	}
}
