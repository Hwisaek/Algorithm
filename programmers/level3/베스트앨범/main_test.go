package main

import (
	"reflect"
	"sort"
	"testing"
)

type song struct {
	id   int
	play int
}
type genreCount struct {
	name  string
	count int
}

func solution(genres []string, plays []int) (album []int) {
	genreCategorized := make(map[string][]song)
	genrePlayed := make(map[string]int)

	for i := range plays {
		genre, play := genres[i], plays[i]
		genreCategorized[genre] = append(genreCategorized[genre], song{id: i, play: play})
		genrePlayed[genre] += play
	}

	genreRank := make([]genreCount, 0, len(genrePlayed))
	for name, count := range genrePlayed {
		genreRank = append(genreRank, genreCount{name: name, count: count})
	}

	for _, v := range genreCategorized {
		sort.Slice(v, func(i, j int) bool {
			if v[i].play == v[j].play {
				return v[i].id < v[j].id
			}
			return v[i].play > v[j].play
		})
	}
	sort.Slice(genreRank, func(i, j int) bool {
		return genreRank[i].count > genreRank[j].count
	})

	for _, g := range genreRank {
		for i, s := range genreCategorized[g.name] {
			if i > 1 {
				break
			}
			album = append(album, s.id)
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
