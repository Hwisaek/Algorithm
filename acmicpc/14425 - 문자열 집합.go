package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, m int
	_, _ = fmt.Fscanln(rd, &n, &m)

	set := make(map[string]bool)
	for i := 0; i < n; i++ {
		str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
		set[str] = true
	}

	result := 0
	for i := 0; i < m; i++ {
		str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
		if set[str] {
			result++
		}
	}

	_, _ = wr.WriteString(fmt.Sprintf("%d", result))
	_ = wr.Flush()
}
