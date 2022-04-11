package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func main() {
	var N int
	_, _ = fmt.Scanln(&N)
	fmt.Println(solution1436(N))
}

func solution1436(n int) (result int) {
	result = 666
	num := 1666
	for n > 1 {
		if strings.Count(strconv.Itoa(num), "666") > 0 {
			n--
			result = num
		}
		num++
	}
	return
}

func Benchmark1436(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution1436(i)
	}
}

func Test_solution1436(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{name: "", args: args{n: 2}, wantResult: 1666},
		{name: "", args: args{n: 3}, wantResult: 2666},
		{name: "", args: args{n: 6}, wantResult: 5666},
		{name: "", args: args{n: 187}, wantResult: 66666},
		{name: "", args: args{n: 500}, wantResult: 166699},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution1436(tt.args.n); gotResult != tt.wantResult {
				t.Errorf("solution1436() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
