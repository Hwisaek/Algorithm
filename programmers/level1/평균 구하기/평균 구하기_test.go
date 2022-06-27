package 평균_구하기

import "testing"

func solution(arr []int) float64 {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return float64(sum) / float64(len(arr))
}

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "",
			args: args{[]int{1, 2, 3, 4}},
			want: 2.5,
		},
		{
			name: "",
			args: args{[]int{5, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
