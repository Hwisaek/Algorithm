package 하샤드_수

import "testing"

func solution(x int) bool {
	sum := 0
	n := x
	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	if x%sum == 0 {
		return true
	} else {
		return false
	}
}

func Test_solution(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{x: 10},
			want: true,
		},
		{
			name: "",
			args: args{x: 12},
			want: true,
		},
		{
			name: "",
			args: args{x: 11},
			want: false,
		},
		{
			name: "",
			args: args{x: 13},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.x); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
