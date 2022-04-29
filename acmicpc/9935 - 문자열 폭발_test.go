package main

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	str := scan9935(rd)
	boom := scan9935(rd)

	_, _ = wr.WriteString(solution9935(str, boom))
	_ = wr.Flush()
}

func solution9935(str, boom string) string {
	r, _ := regexp.Compile(boom)
	for {
		idx := strings.Index(str, boom)
		if idx == -1 {
			break
		}
		str = r.ReplaceAllString(str, "")
	}
	if str == "" {
		return "FRULA"
	}
	return str
}

func scan9935(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark9935(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution9935("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb", "ab")
	}
}

func Test_solution9935(t *testing.T) {
	type args struct {
		str  string
		boom string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{str: "mirkovC4nizCC44", boom: "C4"}, wantResult: "mirkovniz"},
		{name: "", args: args{str: "12ab112ab2ab", boom: "12ab"}, wantResult: "FRULA"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution9935(tt.args.str, tt.args.boom); gotResult != tt.wantResult {
				t.Errorf("solution9935() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
