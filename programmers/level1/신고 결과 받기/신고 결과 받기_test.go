package main

import (
	"reflect"
	"strings"
	"testing"
)

func solution(idList []string, report []string, k int) (result []int) {
	result = make([]int, len(idList))
	reportMap := make(map[string]map[string]interface{})

	for _, s := range report {
		split := strings.Split(s, " ")
		user, reported := split[0], split[1]

		if _, exist := reportMap[reported]; !exist {
			reportMap[reported] = make(map[string]interface{})
		}
		reportMap[reported][user] = struct{}{}
	}

	resultMap := make(map[string]int)
	for _, reporter := range reportMap {
		if len(reporter) >= k {
			for user := range reporter {
				resultMap[user]++
			}
		}
	}

	for i, s := range idList {
		result[i] = resultMap[s]
	}
	return
}

func Test_solution(t *testing.T) {
	type args struct {
		idList []string
		report []string
		k      int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "", args: args{
			idList: []string{"muzi", "frodo", "apeach", "neo"},
			report: []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
			k:      2,
		}, want: []int{2, 1, 1, 0}},
		{name: "", args: args{
			idList: []string{"muzi", "frodo", "apeach", "neo"},
			report: []string{"muzi frodo", "muzi frodo", "muzi frodo", "muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"},
			k:      2,
		}, want: []int{2, 1, 1, 0}},
		{name: "", args: args{
			idList: []string{"con", "ryan"},
			report: []string{"ryan con", "ryan con", "ryan con", "ryan con"},
			k:      3,
		}, want: []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solution(tt.args.idList, tt.args.report, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solution() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkSolution(b *testing.B) {
	for i := 0; i < b.N; i++ {
		solution([]string{"muzi", "frodo", "apeach", "neo"}, []string{"muzi frodo", "apeach frodo", "frodo neo", "muzi neo", "apeach muzi"}, 2)
	}
}
