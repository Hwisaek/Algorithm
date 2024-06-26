package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	r.Scan()
	n, _ := strconv.Atoi(r.Text())

	x := list.New()
	for i := 0; i < n; i++ {
		r.Scan()
		s := strings.Split(r.Text(), " ")
		switch s[0] {
		case "1":
			y, _ := strconv.Atoi(s[1])
			x.PushFront(y)
		case "2":
			y, _ := strconv.Atoi(s[1])
			x.PushBack(y)
		case "3":
			if x.Len() > 0 {
				y := x.Front()
				x.Remove(y)
				w.WriteString(fmt.Sprintln(y.Value))
			} else {
				w.WriteString("-1\n")
			}
		case "4":
			if x.Len() > 0 {
				y := x.Back()
				x.Remove(y)
				w.WriteString(fmt.Sprintln(y.Value))
			} else {
				w.WriteString("-1\n")
			}
		case "5":
			w.WriteString(fmt.Sprintln(x.Len()))
		case "6":
			if l := x.Len(); l == 0 {
				w.WriteString("1\n")
			} else {
				w.WriteString("0\n")
			}
		case "7":
			if x.Len() > 0 {
				y := x.Front()
				w.WriteString(fmt.Sprintln(y.Value))
			} else {
				w.WriteString("-1\n")
			}
		case "8":
			if x.Len() > 0 {
				y := x.Back()
				w.WriteString(fmt.Sprintln(y.Value))
			} else {
				w.WriteString("-1\n")
			}
		}
	}
}
