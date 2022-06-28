package 이진_변환_반복하기

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func solution(s string) (result []int) {
	result = make([]int, 2)
	for {
		result[1] += strings.Count(s, "0")
		s = strings.ReplaceAll(s, "0", "")

		s = strconv.FormatInt(int64(len(s)), 2)
		result[0]++

		if s == "1" {
			break
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
		want []int
	}{
		{name: "", args: args{s: "110010101001"}, want: []int{3, 8}},
		{name: "", args: args{s: "01110"}, want: []int{3, 3}},
		{name: "", args: args{s: "1111111"}, want: []int{4, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
