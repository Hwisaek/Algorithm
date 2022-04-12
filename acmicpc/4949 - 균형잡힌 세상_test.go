package main

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		scanner.Scan()
		str := scanner.Text()
		if str == "." {
			break
		}
		fmt.Println(solution4949(str))
	}
}

func solution4949(str string) (result string) {
	stack := []string{}
	isBracket := map[string]bool{
		"[": true,
		"]": true,
		"(": true,
		")": true,
		"{": true,
		"}": true,
	}
	pairBracket := map[string]string{
		"[": "]",
		"]": "[",
		"(": ")",
		")": "(",
		"{": "}",
		"}": "{",
	}

	for _, i32 := range str {
		s := string(i32)
		if isBracket[s] {
			stack = append(stack, s)
		}
	}
	if len(stack) == 0 {
		return "yes"
	}
	if stack[0] == ")" || stack[0] == "}" || stack[0] == "]" {
		return "no"
	}
	if len(stack) > 0 && (stack[len(stack)-1] == "(" || stack[len(stack)-1] == "{" || stack[len(stack)-1] == "[") {
		return "no"
	}

	check := []string{}
	for _, s := range stack {
		if s == ")" || s == "}" || s == "]" {
			if len(check) == 0 {
				return "no"
			}
			if pairBracket[s] == check[len(check)-1] {
				check = check[:len(check)-1]
			} else {
				return "no"
			}
		} else {
			check = append(check, s)
		}
	}
	if len(check) > 0 {
		return "no"
	}

	return "yes"
}

func Benchmark4949(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution4949("So when I die (the [first] I will see in (heaven) is a score list).")
	}
}

func Test_solution4949(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name       string
		args       args
		wantResult string
	}{
		{name: "", args: args{str: "So when I die (the [first] I will see in (heaven) is a score list)."}, wantResult: "yes"},
		{name: "", args: args{str: "[ first in ] ( first out )."}, wantResult: "yes"},
		{name: "", args: args{str: "Half Moon tonight (At least it is better than no Moon at all]."}, wantResult: "no"},
		{name: "", args: args{str: "A rope may form )( a trail in a maze."}, wantResult: "no"},
		{name: "", args: args{str: "Help( I[m being held prisoner in a fortune cookie factory)].\n"}, wantResult: "no"},
		{name: "", args: args{str: "([ (([( [ ] ) ( ) (( ))] )) ])."}, wantResult: "yes"},
		{name: "", args: args{str: " ."}, wantResult: "yes"},
		{name: "", args: args{str: "("}, wantResult: "no"},
		{name: "", args: args{str: "["}, wantResult: "no"},
		{name: "", args: args{str: ")"}, wantResult: "no"},
		{name: "", args: args{str: "]"}, wantResult: "no"},
		{name: "", args: args{str: "]"}, wantResult: "no"},
		{name: "", args: args{str: "(((("}, wantResult: "no"},
		{name: "", args: args{str: "([)"}, wantResult: "no"},
		{name: "", args: args{str: ".."}, wantResult: "yes"},
		{name: "", args: args{str: "([)]"}, wantResult: "no"},
		{name: "", args: args{str: "(])]"}, wantResult: "no"},
		{name: "", args: args{str: "())"}, wantResult: "no"},
		{name: "", args: args{str: "a"}, wantResult: "yes"},
		{name: "", args: args{str: "(( [[ ]) ])"}, wantResult: "no"},
		{name: "", args: args{str: "[(()])."}, wantResult: "no"},
		{name: "", args: args{str: "t [least it is b[etter than no Moon at all]"}, wantResult: "no"},
		{name: "", args: args{str: "[least it is b[etter than no Moon at all]"}, wantResult: "no"},
		{name: "", args: args{str: "Half Moon tonight (At least it is better than no Moon at all]."}, wantResult: "no"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := solution4949(tt.args.str); gotResult != tt.wantResult {
				t.Errorf("solution4949() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
