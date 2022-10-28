package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
)

func main() {
	input, _ := r.ReadString('\n')

	_, _ = fmt.Fprintln(w, len(regexp.MustCompile("c=|c-|dz=|d-|lj|nj|s=|z=|[a-z]").FindAllString(strings.TrimSpace(input), -1)))
	_ = w.Flush()
}
