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

	word := scan10808(rd)

	_, _ = wr.WriteString(solution10808(word))
	_ = wr.Flush()
}

func solution10808(word string) (result string) {
	count := make([]int, 26)
	for _, c := range word {
		char := rune(c) - rune('a')
		count[char]++
	}

	for _, i := range count {
		result += fmt.Sprintf("%d ", i)
	}
	result = strings.TrimSpace(result)
	return
}

func scan10808(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark10808(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution10808("onetwothreefourfivesixseveneightninetenonetwothreefourfivesixseveneightninetenonetwothreefourfivesix")
	}
}

func Test_solution10808(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{word: "baekjoon"}, wantResult: "1 1 0 0 1 0 0 0 0 1 1 0 0 1 2 0 0 0 0 0 0 0 0 0 0 0"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution10808(tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution10808() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
