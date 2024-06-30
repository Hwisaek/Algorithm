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

	T, _ := strconv.Atoi(scan())

	for i := 0; i < T; i++ {
		s := scan()
		r, c := isPalindrome(s)
		fmt.Fprintln(w, r, c)
	}
}

func recursion(s string, l, r, c int) (int, int) {
	c++
	if l >= r {
		return 1, c
	} else if s[l] != s[r] {
		return 0, c
	} else {
		return recursion(s, l+1, r-1, c)
	}
}

func isPalindrome(s string) (int, int) {
	return recursion(s, 0, len(s)-1, 0)
}
