package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, m int
	_, _ = fmt.Fscanln(rd, &n, &m)

	set := make(map[string]bool)
	for i := 0; i < n; i++ {
		str := scan14425(rd)
		set[str] = true
	}

	result := 0
	for i := 0; i < m; i++ {
		str := scan14425(rd)
		if set[str] {
			result++
		}
	}

	_, _ = fmt.Fprint(wr, result)
	_ = wr.Flush()
}

func scan14425(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n')
	str = strings.TrimSpace(str) // 마지막 줄바꿈 문자가 포함되므로 제거
	return str
}
