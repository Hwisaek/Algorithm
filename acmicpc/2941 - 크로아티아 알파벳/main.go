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
	regex := regexp.MustCompile("c=|c-|dz=|d-|lj|nj|s=|z=|[a-z]")
	alphabets := regex.FindAllString(strings.TrimSpace(input), -1)
	_, _ = fmt.Fprintln(w, len(alphabets))
	_ = w.Flush()
}
