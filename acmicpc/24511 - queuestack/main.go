package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	r.ReadString('\n')

	s, _ := r.ReadString('\n')
	A := strings.Fields(s)

	s, _ = r.ReadString('\n')
	B := strings.Fields(s)

	l := list.New()
	for i := 0; i < len(B); i++ {
		if A[i] == "0" {
			l.PushBack(B[i])
		}
	}

	r.ReadString('\n')

	s, _ = r.ReadString('\n')
	C := strings.Fields(s)
	for _, c := range C {
		l.PushFront(c)
		x := l.Back()
		fmt.Fprint(w, x.Value, " ")
		l.Remove(x)
	}
}
