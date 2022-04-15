package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	n, _ := strconv.Atoi(str)

	wr.WriteString(fmt.Sprintf(`%d`, solution2164(n)))
}

func solution2164(n int) (result int) {
	queue := []int{}
	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}

	for len(queue) > 1 {
		queue = queue[1:]
		queue = append(queue[1:], queue[0])
	}
	return queue[0]
}

func Benchmark2164(b *testing.B) {
	for i := 1; i < b.N; i++ {
		solution2164(i)
	}
}

func Test_solution2164(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{name: "", args: args{n: 6}, wantResult: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution2164(tt.args.n); gotResult != tt.wantResult {
				t.Errorf("solution2164() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
