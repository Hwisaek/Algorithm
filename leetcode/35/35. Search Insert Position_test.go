package _5

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			wantResult: 2,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			wantResult: 1,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			wantResult: 4,
		},
		{
			name: "",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 0,
			},
			wantResult: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := searchInsert(tt.args.nums, tt.args.target); gotResult != tt.wantResult {
				t.Errorf("searchInsert() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
