package JadenCase_문자열_만들기

import (
	"strings"
	"testing"
	"unicode"
)

func solution(s string) (result string) {
	result = ""
	words := strings.Split(s, " ")
	for i, word := range words {
		for j, c := range word {
			if unicode.IsLetter(c) {
				if j == 0 {
					result += strings.ToUpper(string(c))
				} else {
					result += strings.ToLower(string(c))
				}
			} else {
				result += string(c)
			}
		}
		if i != len(words)-1 {
			result += " "
		}
	}
	return
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
		{name: "", args: args{s: "3people unFollowed me"}, want: "3people Unfollowed Me"},
		{name: "", args: args{s: "for the last week"}, want: "For The Last Week"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
