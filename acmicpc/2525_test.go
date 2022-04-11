package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	var A, B, C int
	_, _ = fmt.Fscanln(reader, &A, &B)
	_, _ = fmt.Fscanln(reader, &C)
	fmt.Println(solution(A, B, C))
}

func solution(hour, min, duration int) (result string) {
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

func Test_solution(t *testing.T) {
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
			if gotResult := solution(tt.args.hour, tt.args.min, tt.args.duration); gotResult != tt.wantResult {
				t.Errorf("solution() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
