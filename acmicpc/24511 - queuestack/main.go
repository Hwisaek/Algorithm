package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	sc.Scan()

	sc.Scan()
	A := strings.Split(sc.Text(), " ")

	sc.Scan()
	B := strings.Split(sc.Text(), " ")

	if len(A) != len(B) {
		panic("test")
	}

	l := list.New()
	for i := 0; i < len(B); i++ {
		if A[i] == "0" {
			l.PushBack(B[i])
		}
	}

	sc.Scan()

	sc.Scan()
	for _, s := range strings.Split(sc.Text(), " ") {
		l.PushFront(s)
		x := l.Back()
		fmt.Fprint(w, x.Value, " ")
		l.Remove(x)
	}
}
