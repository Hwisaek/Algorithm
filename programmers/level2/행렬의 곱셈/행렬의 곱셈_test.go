package 행렬의_곱셈

import (
	"reflect"
	"testing"
)

func solution(arr1 [][]int, arr2 [][]int) (result [][]int) {
	result = make([][]int, len(arr1))
	for i := 0; i < len(arr1); i++ {
		result[i] = make([]int, len(arr2[0]))
		for j := 0; j < len(arr2[0]); j++ {
			for k := 0; k < len(arr2); k++ {
				result[i][j] += arr1[i][k] * arr2[k][j]
			}
		}
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		arr1 [][]int
		arr2 [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				arr1: [][]int{{1, 4}, {3, 2}, {4, 1}},
				arr2: [][]int{{3, 3}, {3, 3}},
			},
			want: [][]int{{15, 15}, {15, 15}, {15, 15}},
		},
		{
			name: "",
			args: args{
				arr1: [][]int{{2, 3, 2}, {4, 2, 4}, {3, 1, 4}},
				arr2: [][]int{{5, 4, 3}, {2, 4, 1}, {3, 1, 1}},
			},
			want: [][]int{{22, 22, 11}, {36, 28, 18}, {29, 20, 14}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
