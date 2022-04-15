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

	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	str_arr := strings.Split(str, " ")
	n, _ := strconv.Atoi(str_arr[0])
	k, _ := strconv.Atoi(str_arr[1])

	_, _ = wr.WriteString(solution11866(n, k))
	_ = wr.Flush()
}

func solution11866(n, k int) (result string) {
	result = "<"
	queue := []int{}
	for i := 0; i < n; i++ {
		queue = append(queue, i+1)
	}
	size := len(queue)
	for size != 0 {
		idx := (k - 1) % size
		result += fmt.Sprintf(`%d, `, queue[idx])
		queue = append(queue[idx+1:], queue[:idx]...)
		size = len(queue)
	}
	result = result[:len(result)-2] + ">"
	return
}

func Benchmark11866(b *testing.B) {
	for i := 1; i < b.N; i++ {
		solution11866(1000, 1000)
	}
}

func Test_solution11866(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{
			n: 7,
			k: 3,
		}, wantResult: "<3, 6, 2, 7, 5, 1, 4>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution11866(tt.args.n, tt.args.k); gotResult != tt.wantResult {
				t.Errorf("solution11866() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
