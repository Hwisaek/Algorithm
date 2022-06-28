package k진수에서_소수_개수_구하기

import (
	"regexp"
	"strconv"
	"testing"
)

func solution(n int, k int) (result int) {
	kNum := strconv.FormatInt(int64(n), k)

	regex, _ := regexp.Compile("0+")
	words := regex.Split(kNum, -1)
	for _, word := range words {
		num, _ := strconv.Atoi(word)
		if isPrime(num) {
			result++
		}
	}
	return
}

func isPrime(n int) bool {
	if n == 1 || n == 0 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Test_solution(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				n: 437674,
				k: 3,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				n: 110011,
				k: 10,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
