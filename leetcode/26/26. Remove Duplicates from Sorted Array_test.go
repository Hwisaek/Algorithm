package _26

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "",
			args: args{
				nums: []int{1, 1, 2},
			},
			wantResult: 2,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			wantResult: 5,
		},
		{
			name: "",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			wantResult: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := removeDuplicates(tt.args.nums); gotResult != tt.wantResult {
				t.Errorf("removeDuplicates() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
