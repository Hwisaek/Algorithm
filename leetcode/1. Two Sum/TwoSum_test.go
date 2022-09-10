package ___Two_Sum

import (
	"reflect"
	"testing"
)

func twoSum(nums []int, target int) (result []int) {
	m := map[int]int{}

	for i, num := range nums {
		j, ok := m[target-num]
		if ok {
			return []int{i, j}
		}
		m[num] = i
	}
	return
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name       string
		args       args
		wantResult []int
	}{
		{
			name: "", args: args{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			wantResult: []int{0, 1},
		},
		{
			name: "", args: args{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			wantResult: []int{1, 2},
		},
		{
			name: "", args: args{
				nums:   []int{3, 3},
				target: 6,
			},
			wantResult: []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("twoSum() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
