package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strings"
	"testing"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	word := scan1439(rd)

	_, _ = wr.WriteString(solution1439(word))
	_ = wr.Flush()
}

func solution1439(word string) (result string) {
	rgxZero, _ := regexp.Compile("0+")
	zeros := rgxZero.FindAllString(word, -1)
	rgxOne, _ := regexp.Compile("1+")
	ones := rgxOne.FindAllString(word, -1)

	x := float64(len(zeros))
	y := float64(len(ones))
	result = fmt.Sprintf("%.0f", math.Min(x, y))
	return
}

func scan1439(rd *bufio.Reader) string {
	str, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
	return str
}

func Benchmark1439(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1439("101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101")
	}
}

func Test_solution1439(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{word: "0001100"}, wantResult: "1"},
		{name: "", args: args{word: "11111"}, wantResult: "0"},
		{name: "", args: args{word: "00000001"}, wantResult: "1"},
		{name: "", args: args{word: "11001100110011000001"}, wantResult: "4"},
		{name: "", args: args{word: "11101101"}, wantResult: "2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1439(tt.args.word); gotResult != tt.wantResult {
				t.Errorf("solution1439() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
