package main

import (
	"fmt"
	"math"
	"testing"
)

func main() {
	var a, b, c int
	_, _ = fmt.Scanln(&a, &b, &c)
	fmt.Println(solution2480(a, b, c))
}

func solution2480(a, b, c int) (result int) {
	if a == b && b == c {
		result = 10000 + a*1000
	} else if a == b || b == c || a == c {
		var n int
		if a == b {
			n = a
		} else if b == c {
			n = b
		} else {
			n = c
		}
		result = 1000 + n*100
	} else {
		max := int(math.Max(float64(a), math.Max(float64(b), float64(c))))
		result = max * 100
	}
	return
}

func Test_solution2480(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			name: "1",
			args: args{
				a: 3,
				b: 3,
				c: 6,
			},
			wantResult: 1300,
		},
		{
			name: "2",
			args: args{
				a: 2,
				b: 2,
				c: 2,
			},
			wantResult: 12000,
		},
		{
			name: "3",
			args: args{
				a: 6,
				b: 2,
				c: 5,
			},
			wantResult: 600,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution2480(tt.args.a, tt.args.b, tt.args.c); gotResult != tt.wantResult {
				t.Errorf("solution() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
