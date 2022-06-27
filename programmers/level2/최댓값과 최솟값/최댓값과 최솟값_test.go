package 최댓값과_최솟값

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
)

func solution(s string) (result string) {
	words := strings.Split(s, " ")

	first, _ := strconv.Atoi(words[0])
	min := float64(first)
	max := float64(first)

	for _, word := range words {
		num, _ := strconv.Atoi(word)
		min = math.Min(min, float64(num))
		max = math.Max(max, float64(num))
	}
	return fmt.Sprintf("%.f %.f", min, max)
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
		{name: "", args: args{s: "1 2 3 4"}, want: "1 4"},
		{name: "", args: args{s: "-1 -2 -3 -4"}, want: "-4 -1"},
		{name: "", args: args{s: "-1 -1"}, want: "-1 -1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
