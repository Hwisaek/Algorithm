package 이상한_문자_만들기

import (
	"strings"
	"testing"
	"unicode"
)

func solution(s string) (result string) {
	words := strings.Split(s, " ")
	for i, word := range words {
		runes := []rune(word)
		for j, c := range runes {
			if j%2 == 0 {
				runes[j] = unicode.ToUpper(c)
			} else {
				runes[j] = unicode.ToLower(c)
			}
		}
		words[i] = string(runes)
	}

	result = strings.Join(words, " ")
	return
}

func FuzzSolution(f *testing.F) {
	f.Fuzz(func(t *testing.T, orig string) {
		result := solution(orig)
		t.Log(result)
	})
}

func Test_solution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{s: "try hello world"}, want: "TrY HeLlO WoRlD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
