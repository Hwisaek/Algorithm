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

	_, _ = rd.ReadString('\n')
	dict := make(map[int]int)
	input, _ := rd.ReadString('\n')
	strArr := strings.Split(strings.TrimSpace(input), " ")
	for _, s := range strArr {
		num, _ := strconv.Atoi(s)
		dict[num]++
	}

	_, _ = rd.ReadString('\n')
	input, _ = rd.ReadString('\n')
	strArr = strings.Split(strings.TrimSpace(input), " ")
	for _, s := range strArr {
		num, _ := strconv.Atoi(s)
		_, _ = fmt.Fprint(wr, dict[num], " ")
	}

	_ = wr.Flush()
}
