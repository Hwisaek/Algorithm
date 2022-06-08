package 멀쩡한_사각형

import (
	"testing"
)

func solution(w int, h int) (answer int64) {
	w64, h64 := int64(w), int64(h)
	answer = w64*h64 - (w64 + h64 - gcd(w64, h64))
	return
}

func gcd(w int64, h int64) int64 {
	var result int64
	for i := int64(1); i <= w && i <= h; i++ {
		if w%i == 0 && h%i == 0 {
			result = i
		}
	}
	return result
}

func Test_solution(t *testing.T) {
	type args struct {
		w int
		h int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			name: "",
			args: args{
				w: 8,
				h: 12,
			},
			want: 80,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.w, tt.args.h); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
