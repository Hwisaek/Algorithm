package 단어_변환

import (
	"testing"
)

func solution(begin string, target string, words []string) (count int) {
	if !isContains(words, target) {
		return 0
	}

	isVisited := make([]bool, len(words))
	type q struct {
		str   string
		count int
	}
	now := q{
		str:   begin,
		count: 0,
	}
	queue := []q{now}
	for queue[0].str != target {
		now = queue[0]
		queue = queue[1:]
		for i, word := range words {
			if !isVisited[i] && isChangeable(now.str, word) {
				queue = append(queue, q{
					str:   word,
					count: now.count + 1,
				})
				isVisited[i] = true
			}
		}
	}

	count = queue[0].count
	return
}

func isContains(words []string, target string) (flag bool) {
	for _, word := range words {
		if target == word {
			flag = true
			break
		}
	}
	return flag
}

func isChangeable(begin, target string) (changeable bool) {
	count := 0
	for i := range begin {
		if count > 1 {
			return false
		}
		if begin[i] != target[i] {
			count++
		}
	}
	return count == 1
}

func Test_solution(t *testing.T) {
	type args struct {
		begin  string
		target string
		words  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "", args: args{
				begin:  "hit",
				target: "cog",
				words:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 4,
		},
		{
			name: "",
			args: args{
				begin:  "hit",
				target: "cog",
				words:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.begin, tt.args.target, tt.args.words); got != tt.want {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
