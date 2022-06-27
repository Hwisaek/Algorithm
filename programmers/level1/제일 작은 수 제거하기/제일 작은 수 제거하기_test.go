package 제일_작은_수_제거하기

import (
	"reflect"
	"sort"
	"testing"
)

func solution(arr []int) (result []int) {
	if len(arr) == 1 {
		return []int{-1}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})

	return arr[:len(arr)-1]
}

func Test_solution(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{[]int{4, 3, 2, 1}},
			want: []int{4, 3, 2},
		},
		{
			name: "",
			args: args{[]int{10}},
			want: []int{-1},
		},
		{
			name: "",
			args: args{[]int{1, 2, 3, 4, 5}},
			want: []int{5, 4, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
