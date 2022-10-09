package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
	"strings"
)

var (
	r   = bufio.NewReader(os.Stdin)
	min = -1000
)

func main() {
	_, _ = r.ReadString('\n')
	s, _ := r.ReadString('\n')
	fs := strings.Fields(s)
	arr := make([]int, len(fs))
	for i, f := range fs {
		arr[i], _ = strconv.Atoi(f)
	}

	result := arr[0]
	for i := 1; i < len(arr); i++ {
		arr[i] = max(arr[i-1]+arr[i], arr[i])
		result = max(result, arr[i])
	}

	Println(result)
}

func max(ints ...int) (max int) {
	max = ints[0]
	for _, n := range ints[1:] {
		if n > max {
			max = n
		}
	}
	return
}
