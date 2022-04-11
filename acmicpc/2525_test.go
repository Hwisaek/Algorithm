package main

import (
	"fmt"
	"testing"
)

func main() {
	var A, B, C int
	_, _ = fmt.Scanln(&A, &B)
	_, _ = fmt.Scanln(&C)
	fmt.Println(solution2525(A, B, C))
}

func solution2525(hour, min, duration int) (result string) {
	min += duration
	if min > 59 {
		hour += min / 60
		min %= 60
	}
	if hour > 23 {
		hour -= 24
	}
	return fmt.Sprintf(`%d %d`, hour, min)
}

func Test_solution2525(t *testing.T) {
	type args struct {
		hour     int
		min      int
		duration int
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{
			name: "0",
			args: args{
				hour:     14,
				min:      30,
				duration: 20,
			},
			wantResult: "14 50",
		},
		{
			name: "1",
			args: args{
				hour:     17,
				min:      40,
				duration: 80,
			},
			wantResult: "19 0",
		},
		{
			name: "2",
			args: args{
				hour:     23,
				min:      48,
				duration: 25,
			},
			wantResult: "0 13",
		},
		{
			name: "3",
			args: args{
				hour:     22,
				min:      0,
				duration: 61,
			},
			wantResult: "23 1",
		},
		{
			name: "4",
			args: args{
				hour:     22,
				min:      59,
				duration: 61,
			},
			wantResult: "0 0",
		},
		{
			name: "5",
			args: args{
				hour:     23,
				min:      59,
				duration: 61,
			},
			wantResult: "1 0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution2525(tt.args.hour, tt.args.min, tt.args.duration); gotResult != tt.wantResult {
				t.Errorf("solution2525() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
