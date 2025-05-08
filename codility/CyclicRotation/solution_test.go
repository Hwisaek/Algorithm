package solution

import (
	"reflect"
	"testing"
)

func Solution(A []int, K int) (answer []int) {
	answer = make([]int, 0, len(A))

	if len(A) == 0 {
		return
	}

	l := len(A)
	K = K % l
	for i := 0; i < l; i++ {
		index := i - K
		if index < 0 {
			index += l * ((-index / l) + 1)
		}
		index = index % l
		answer = append(answer, A[index])
	}

	return
}

func TestSolution(t *testing.T) {
	type args struct {
		A []int
		K int
	}

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"test 1", args{[]int{3, 8, 9, 7, 6}, 3}, []int{9, 7, 6, 3, 8}},
		{"test 2", args{[]int{3, 8, 9, 7, 6}, -7}, []int{7, 6, 3, 8, 9}},
		{"test 3", args{[]int{3, 8, 9, 7, 6}, -6}, []int{9, 7, 6, 3, 8}},
		{"test 4", args{[]int{3, 8, 9, 7, 6}, -5}, []int{3, 8, 9, 7, 6}},
		{"test 5", args{[]int{3, 8, 9, 7, 6}, -4}, []int{6, 3, 8, 9, 7}},
		{"test 6", args{[]int{3, 8, 9, 7, 6}, -3}, []int{7, 6, 3, 8, 9}},
		{"test 7", args{[]int{3, 8, 9, 7, 6}, -2}, []int{8, 9, 7, 6, 3}},
		{"test 8", args{[]int{3, 8, 9, 7, 6}, -1}, []int{9, 7, 6, 3, 8}},
		{"test 9", args{[]int{3, 8, 9, 7, 6}, 0}, []int{3, 8, 9, 7, 6}},
		{"test 10", args{[]int{3, 8, 9, 7, 6}, 1}, []int{6, 3, 8, 9, 7}},
		{"test 11", args{[]int{3, 8, 9, 7, 6}, 2}, []int{7, 6, 3, 8, 9}},
		{"test 12", args{[]int{3, 8, 9, 7, 6}, 3}, []int{9, 7, 6, 3, 8}},
		{"test 13", args{[]int{3, 8, 9, 7, 6}, 4}, []int{8, 9, 7, 6, 3}},
		{"test 14", args{[]int{3, 8, 9, 7, 6}, 5}, []int{3, 8, 9, 7, 6}},
		{"test 15", args{[]int{3, 8, 9, 7, 6}, 6}, []int{6, 3, 8, 9, 7}},
		{"test 16", args{[]int{3, 8, 9, 7, 6}, 7}, []int{7, 6, 3, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Solution(tt.args.A, tt.args.K)
			if !reflect.DeepEqual(result, tt.want) {
				t.Errorf("expected %v, got %v", tt.want, result)
			}
		})
	}
}
