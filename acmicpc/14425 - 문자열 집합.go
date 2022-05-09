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
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}
