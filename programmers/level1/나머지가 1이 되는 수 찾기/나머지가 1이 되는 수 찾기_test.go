package main

import "testing"

func solution(n int) (result int) {
	result = n - 1
	for i := 2; i < n; i++ {
		if n%i == 1 {
			return i
		}
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{n: 10}, want: 3},
		{name: "", args: args{n: 12}, want: 11},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(10)
	}
}
