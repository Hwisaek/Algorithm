package main

import (
	"bufio"
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

	now := 1
	s := make([]int, 0, n)

	r.Scan()
	split := strings.Split(r.Text(), " ")
	for i := 0; true; i++ {
		for len(s) > 0 && s[len(s)-1] == now {
			s = s[:len(s)-1]
			now++
		}

		if i >= len(split) {
			break
		}

		j, _ := strconv.Atoi(split[i])
		if j == now {
			now++
			continue
		}

		if len(s) > 0 && s[len(s)-1] < j {
			break
		}
		s = append(s, j)
	}

	if len(s) > 0 {
		fmt.Fprint(w, "Sad")
		return
	}
	fmt.Fprint(w, "Nice")
}
