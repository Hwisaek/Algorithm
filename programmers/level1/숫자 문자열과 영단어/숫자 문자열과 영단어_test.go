package 숫자_문자열과_영단어

import (
	"strconv"
	"strings"
	"testing"
)

func solution(s string) int {
	dict := map[string]string{
		"zero": "0", "one": "1", "two": "2", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
	}
	for key, value := range dict {
		s = strings.ReplaceAll(s, key, value)
	}
	atoi, _ := strconv.Atoi(s)
	return atoi
}

func Test_solution(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{s: "one4seveneight"}, want: 1478},
		{name: "", args: args{s: "23four5six7"}, want: 234567},
		{name: "", args: args{s: "2three45sixseven"}, want: 234567},
		{name: "", args: args{s: "123"}, want: 123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
