package main

import (
	"reflect"
	"testing"
)

func solution(lottos []int, winNums []int) []int {
	count := 0
	joker := 0
	for i := 0; i < len(lottos); i++ {
		if lottos[i] == 0 {
			joker++
		} else if indexOf(winNums, lottos[i]) != -1 {
			count++
		}
	}

	joker += count

	max := 7 - joker
	if max > 6 {
		max = 6
	}
	min := 7 - count
	if min > 6 {
		min = 6
	}
	return []int{max, min}
}

func indexOf(array []int, target int) int {
	for i, num := range array {
		if num == target {
			return i
		}
	}
	return -1
}

func Test_solution(t *testing.T) {
	type args struct {
		lottos  []int
		winNums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				lottos:  []int{44, 1, 0, 0, 31, 25},
				winNums: []int{31, 10, 45, 1, 6, 19},
			},
			want: []int{3, 5},
		},
		{
			name: "",
			args: args{
				lottos:  []int{0, 0, 0, 0, 0, 0},
				winNums: []int{38, 19, 20, 40, 15, 25},
			},
			want: []int{1, 6},
		},
		{
			name: "",
			args: args{
				lottos:  []int{45, 4, 35, 20, 3, 9},
				winNums: []int{20, 9, 3, 45, 4, 35},
			},
			want: []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.lottos, tt.args.winNums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
