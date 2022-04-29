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

	word := scan2743(rd)

	_, _ = wr.WriteString(solution2743(word))
	_ = wr.Flush()
}

func solution2743(word string) (result string) {
	result = fmt.Sprintf("%d", len(word))
	return
}

func scan2743(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark2743(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution2743("pulljima")
	}
}

func Test_solution2743(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{word: "pulljima"}, wantResult: "8"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution2743(tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution2743() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
