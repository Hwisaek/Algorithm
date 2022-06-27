package 콜라츠_추측

import "testing"

func solution(num int) (result int) {
	for num != 1 {
		switch {
		case num%2 == 0:
			num = num / 2
		default:
			num = num*3 + 1
		}
		result++

		if result >= 500 {
			return -1
		}
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{num: 6},
			want: 8,
		},
		{
			name: "",
			args: args{num: 16},
			want: 4,
		},
		{
			name: "",
			args: args{num: 626331},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.num); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
