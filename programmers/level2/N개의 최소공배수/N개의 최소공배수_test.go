package N개의_최소공배수

import "testing"

func solution(arr []int) (result int) {
	result = arr[0]
	for i := 1; i < len(arr); i++ {
		result = getLcm(result, arr[i])
	}
	return
}

func getGcd(a, b int) int {
	if b == 0 {
		return a
	}
	return getGcd(b, a%b)
}

func getLcm(a, b int) int {
	return a * b / getGcd(a, b)
}

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{[]int{2, 6, 8, 14}}, want: 168},
		{name: "", args: args{[]int{1, 2, 3}}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
