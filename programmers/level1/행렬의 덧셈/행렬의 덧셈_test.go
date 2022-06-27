package 행렬의_덧셈

import (
	"reflect"
	"testing"
)

// solution return matrix addition
func solution(arr1 [][]int, arr2 [][]int) (result [][]int) {
	for i1, row := range arr1 {
		for i2, val := range row {
			arr2[i1][i2] += val
		}
	}
	return arr2
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
				arr1: [][]int{{1, 2}, {2, 3}},
				arr2: [][]int{{3, 4}, {5, 6}},
			},
			want: [][]int{{4, 6}, {7, 9}},
		},
		{
			name: "",
			args: args{
				arr1: [][]int{{1}, {2}},
				arr2: [][]int{{3}, {4}},
			},
			want: [][]int{{4}, {6}},
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
