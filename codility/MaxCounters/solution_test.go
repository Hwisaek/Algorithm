package solution

import (
	"reflect"
	"testing"
)

func Solution(N int, A []int) (answer []int) {
	answer = make([]int, N)
	MAX := 0
	counter := 0

	for _, v := range A {
		if v <= N {
			answer[v-1]++
			MAX = max(MAX, answer[v-1])
		} else if v == N+1 {
			answer = make([]int, N)
			counter += MAX
			MAX = 0
		}
	}

	for i := 0; i < N; i++ {
		answer[i] += counter
	}

	return
}

func TestSolution(t *testing.T) {
	type arg struct {
		N int
		A []int
	}
	type expect struct {
		answer []int
	}

	tests := []struct {
		name   string
		arg    arg
		expect expect
	}{
		{
			name: "",
			arg: arg{
				N: 5,
				A: []int{3, 4, 4, 6, 1, 4, 4},
			},
			expect: expect{
				answer: []int{3, 2, 2, 4, 2},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{3, 3, 3},
			},
			expect: expect{
				answer: []int{0, 0},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{1, 2, 1, 2},
			},
			expect: expect{
				answer: []int{2, 2},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{3, 1, 2},
			},
			expect: expect{
				answer: []int{1, 1},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{3, 1},
			},
			expect: expect{
				answer: []int{1, 0},
			},
		},
		{
			name: "",
			arg: arg{
				N: 3,
				A: []int{4, 4, 1, 4, 2},
			},
			expect: expect{
				answer: []int{1, 2, 1},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{1, 3, 2},
			},
			expect: expect{
				answer: []int{1, 2},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{1, 3, 1},
			},
			expect: expect{
				answer: []int{2, 1},
			},
		},
		{
			name: "",
			arg: arg{
				N: 1,
				A: []int{1, 2, 1, 2, 1},
			},
			expect: expect{
				answer: []int{3},
			},
		},
		{
			name: "",
			arg: arg{
				N: 2,
				A: []int{1, 3, 2},
			},
			expect: expect{
				answer: []int{1, 2},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Solution(tt.arg.N, tt.arg.A)
			if !reflect.DeepEqual(got, tt.expect.answer) {
				t.Errorf("%s: Solution() = %v, want %v", tt.name, got, tt.expect.answer)
			}
		})
	}
}
