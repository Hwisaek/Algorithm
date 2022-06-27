package 피보나치_수

import "testing"

func solution(n int) int {
	return fibonacci(n)
}

func fibonacci(n int) (result int) {
	cache := make([]int, n+1)
	cache[0] = 0
	cache[1] = 1
	for i := 2; i <= n; i++ {
		cache[i] = cache[i-1]%1234567 + cache[i-2]%1234567
	}
	return cache[n] % 1234567
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
		{name: "", args: args{n: 2}, want: 1},
		{name: "", args: args{n: 3}, want: 2},
		{name: "", args: args{n: 4}, want: 3},
		{name: "", args: args{n: 5}, want: 5},
		{name: "", args: args{n: 6}, want: 8},
		{name: "", args: args{n: 7}, want: 13},
		{name: "", args: args{n: 8}, want: 21},
		{name: "", args: args{n: 9}, want: 34},
		{name: "", args: args{n: 10}, want: 55},
		{name: "", args: args{n: 11}, want: 89},
		{name: "", args: args{n: 12}, want: 144},
		{name: "", args: args{n: 13}, want: 233},
		{name: "", args: args{n: 14}, want: 377},
		{name: "", args: args{n: 15}, want: 610},
		{name: "", args: args{n: 16}, want: 987},
		{name: "", args: args{n: 17}, want: 1597},
		{name: "", args: args{n: 18}, want: 2584},
		{name: "", args: args{n: 19}, want: 4181},
		{name: "", args: args{n: 20}, want: 6765},
		{name: "", args: args{n: 21}, want: 10946},
		{name: "", args: args{n: 22}, want: 17711},
		{name: "", args: args{n: 23}, want: 28657},
		{name: "", args: args{n: 24}, want: 46368},
		{name: "", args: args{n: 25}, want: 75025},
		{name: "", args: args{n: 26}, want: 121393},
		{name: "", args: args{n: 27}, want: 196418},
		{name: "", args: args{n: 28}, want: 317811},
		{name: "", args: args{n: 29}, want: 514229},
		{name: "", args: args{n: 30}, want: 832040},
		{name: "", args: args{n: 31}, want: 111702},
		{name: "", args: args{n: 32}, want: 943742},
		{name: "", args: args{n: 33}, want: 1055444},
		{name: "", args: args{n: 34}, want: 764619},
		{name: "", args: args{n: 35}, want: 585496},
		{name: "", args: args{n: 36}, want: 115548},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
