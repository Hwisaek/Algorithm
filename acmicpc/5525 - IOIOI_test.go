package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n int
	_, _ = fmt.Fscanln(rd, &n)
	var m int
	_, _ = fmt.Fscanln(rd, &m)
	var s string
	_, _ = fmt.Fscanln(rd, &s)

	_, _ = wr.WriteString(solution5525(n, m, s))
	_ = wr.Flush()
}

func solution5525(n int, m int, s string) (result string) {
	p := "I"
	for i := 0; i < n; i++ {
		p += "OI"
	}

	answer := 0
	count := 0

	i := 0
	for i < m-2 {
		subStr := s[i : i+3]
		if subStr == "IOI" {
			i += 2
			count += 1
			if count == n {
				answer += 1
				count -= 1
			}
		} else {
			i += 1
			count = 0
		}
	}
	return fmt.Sprintf("%d", answer)
}

func scan5525(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n')
	str = strings.TrimSpace(str) // 마지막 줄바꿈 문자가 포함되므로 제거
	return str
}

func Benchmark_solution5525(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution5525(2, 13, "OOIOIOIOIIOII")
	}
}

func Test_solution5525(t *testing.T) {
	type args struct {
		n int
		m int
		s string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{
			n: 1,
			m: 13,
			s: "OOIOIOIOIIOII",
		}, wantResult: "4"},
		{name: "", args: args{
			n: 2,
			m: 13,
			s: "OOIOIOIOIIOII",
		}, wantResult: "2"},
		{name: "", args: args{
			n: 1,
			m: 3,
			s: "IIO",
		}, wantResult: "0"},
		{name: "", args: args{
			n: 1,
			m: 4,
			s: "OIOI",
		}, wantResult: "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution5525(tt.args.n, tt.args.m, tt.args.s); gotResult != tt.wantResult {
				t.Errorf("solution5525() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
