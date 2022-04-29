package main

import (
	"bufio"
	"os"
	"sort"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	word := scan11656(rd)

	_, _ = wr.WriteString(solution11656(word))
	_ = wr.Flush()
}

func solution11656(word string) (result string) {
	strArr := make([]string, 0, len(word))
	for i := range word {
		strArr = append(strArr, word[i:])
	}

	sort.Slice(strArr, func(i, j int) bool {
		return strArr[i] < strArr[j]
	})

	result = strings.Join(strArr, "\n")
	return
}

func scan11656(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark11656(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution11656("baekjoon")
	}
}

func Test_solution11656(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{word: "baekjoon"}, wantResult: "aekjoon\nbaekjoon\nekjoon\njoon\nkjoon\nn\non\noon"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution11656(tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution11656() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
