package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	document := scan1543(rd)
	word := scan1543(rd)

	_, _ = wr.WriteString(solution1543(document, word))
	_ = wr.Flush()
}

func solution1543(document, word string) (result string) {
	return strconv.Itoa(strings.Count(document, word))
}

func scan1543(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark1543(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1543("ababababa", "aba")
	}
}

func Test_solution1543(t *testing.T) {
	type args struct {
		document string
		word     string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{document: "ababababa", word: "aba"}, wantResult: "2"},
		{name: "", args: args{document: "a a a a a", word: "a a"}, wantResult: "2"},
		{name: "", args: args{document: "ababababa", word: "ababa"}, wantResult: "1"},
		{name: "", args: args{document: "aaaaaaa", word: "aa"}, wantResult: "3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1543(tt.args.document, tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution1543() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
