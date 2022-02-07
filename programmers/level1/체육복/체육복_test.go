package 체육복

import (
	"sort"
	"testing"
)

func solution(n int, lost []int, reserve []int) (result int) {
	result += n - len(lost)
	sort.Ints(lost)
	sort.Ints(reserve)
	for idx := len(lost) - 1; idx >= 0; idx-- {
		v := lost[idx]
		if indexOf(&reserve, v) {
			lost = append(lost[:idx], lost[idx+1:]...)
			result++
		}
	}
	for _, v := range lost {
		if indexOf(&reserve, v-1) || indexOf(&reserve, v+1) {
			result++
		}
	}
	return
}

func indexOf(reserve *[]int, target int) bool {
	for idx, v := range *reserve {
		if v == target {
			*reserve = append((*reserve)[:idx], (*reserve)[idx+1:]...)
			return true
		} else if v > target {
			return false
		}
	}
	return false
}

func Test_solution(t *testing.T) {
	type args struct {
		n       int
		lost    []int
		reserve []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n:       5,
				lost:    []int{2, 4},
				reserve: []int{1, 3, 5},
			},
			want: 5,
		},
		{
			name: "2",
			args: args{
				n:       5,
				lost:    []int{2, 4},
				reserve: []int{3},
			},
			want: 4,
		},
		{
			name: "3",
			args: args{
				n:       3,
				lost:    []int{3},
				reserve: []int{1},
			},
			want: 2,
		},
		{
			name: "4",
			args: args{
				n:       5,
				lost:    []int{4, 2},
				reserve: []int{2},
			},
			want: 4,
		},
		{
			name: "5",
			args: args{
				n:       10,
				lost:    []int{2, 4, 5, 6, 7},
				reserve: []int{2, 4, 5, 6, 7},
			},
			want: 10,
		},
		{
			name: "6",
			args: args{
				n:       5,
				lost:    []int{4, 3},
				reserve: []int{3, 2},
			},
			want: 4,
		},
		{
			name: "7",
			args: args{
				n:       6,
				lost:    []int{6, 2, 4},
				reserve: []int{1, 5, 3},
			},
			want: 6,
		},
		{
			name: "8",
			args: args{
				n:       30,
				lost:    []int{18, 21},
				reserve: []int{3, 4, 5, 6, 8, 13, 17, 22, 23, 26, 28},
			},
			want: 30,
		},
		{
			name: "9",
			args: args{
				n:       3,
				lost:    []int{1, 2},
				reserve: []int{2, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.n, tt.args.lost, tt.args.reserve); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_solution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution(30, []int{1, 2, 7, 9, 10, 11, 12, 14, 15, 16, 18, 20, 21, 24, 25, 27}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 20, 22, 23, 24, 25, 26, 27, 28})
	}
}
