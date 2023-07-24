package main

import (
	"testing"
)

func Test_solution(t *testing.T) {
	type args struct {
		numbers string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{numbers: "17"},
			want: 3,
		},
		{
			name: "2",
			args: args{numbers: "011"},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.numbers); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
