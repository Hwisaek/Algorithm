package solution

import (
	"strconv"
	"testing"
)

func Solution(N int) (answer int) {
	binary := strconv.FormatInt(int64(N), 2)

	count := 0
	for _, v := range binary {
		if v == '0' {
			count++
		} else {
			answer = max(answer, count)
			count = 0
		}
	}
	return
}

func TestSolution(t *testing.T) {
	type args struct {
		N int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{9}, 2},
		{"", args{529}, 4},
		{"", args{20}, 1},
		{"", args{15}, 0},
		{"", args{32}, 0},
		{"", args{1041}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solution(tt.args.N)
			if result != tt.want {
				t.Errorf("expected %d, got %d", tt.want, result)
			}
		})
	}
}
