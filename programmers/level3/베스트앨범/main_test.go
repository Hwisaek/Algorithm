package main

import (
	"reflect"
	"sort"
	"testing"
)

type song struct {
	idx, play int
}

func solution(genres []string, plays []int) (result []int) {
	songs := make(map[string][]song)
	total := make(map[string]int)

	for idx := range plays {
		genre, play := genres[idx], plays[idx]
		songs[genre] = append(songs[genre], song{idx, play})
		total[genre] += play
	}

	arr := make([][2]interface{}, 0, len(total))
	for k, v := range total {
		arr = append(arr, [2]interface{}{k, v})
	}

	for _, v := range songs {
		sort.Slice(v, func(i, j int) bool {
			if v[i].play == v[j].play {
				return v[i].idx < v[j].idx
			}
			return v[i].play > v[j].play
		})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1].(int) > arr[j][1].(int)
	})

	for _, e := range arr {
		for i, s := range songs[e[0].(string)] {
			if i > 1 {
				break
			}
			result = append(result, s.idx)
		}
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		genres []string
		plays  []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				genres: []string{"classic", "pop", "classic", "classic", "pop"},
				plays:  []int{500, 600, 150, 800, 2500},
			},
			want: []int{4, 1, 3, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.genres, tt.args.plays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}
