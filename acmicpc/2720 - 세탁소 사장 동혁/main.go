package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var N, B int
	Fscan(r, &N, &B)
	Fprint(w, strings.ToUpper(strconv.FormatInt(int64(N), B)))
}
