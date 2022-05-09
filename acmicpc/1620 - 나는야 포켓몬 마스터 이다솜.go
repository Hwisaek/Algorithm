package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, m int
	_, _ = fmt.Fscanln(rd, &n, &m)

	nameMap := make(map[string]int)
	numMap := make(map[int]string)
	for i := 0; i < n; i++ {
		input, _ := rd.ReadString('\n')
		str := strings.TrimSpace(input)
		nameMap[str] = i + 1
		numMap[i+1] = str
	}

	for i := 0; i < m; i++ {
		input, _ := rd.ReadString('\n')
		str := strings.TrimSpace(input)
		num, err := strconv.Atoi(str)
		if err != nil {
			result := nameMap[str]
			_, _ = fmt.Fprintln(wr, result)
		} else {
			result := numMap[num]
			_, _ = fmt.Fprintln(wr, result)
		}
	}

	_ = wr.Flush()
}
