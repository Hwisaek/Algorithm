package main

import (
	"sort"
)

func solution(n int, costs [][]int) (answer int) {
	sort.Slice(costs, func(i, j int) bool {
		return costs[i][2] < costs[j][2]
	})
	linked := map[int]struct{}{}

	linked[costs[0][0]] = struct{}{}
	linked[costs[0][1]] = struct{}{}
	answer += costs[0][2]
	costs = costs[1:]

	for i := 0; len(linked) < n; {
		cost := costs[i]
		_, ok1 := linked[cost[0]]
		_, ok2 := linked[cost[1]]

		if !ok1 && !ok2 {
			i++
			continue
		}

		linked[cost[0]] = struct{}{}
		linked[cost[1]] = struct{}{}

		answer += cost[2]
		costs = append(costs[:i], costs[i+1:]...)
		i = 0
	}

	return
}

func main() {
	solution(4, [][]int{{0, 1, 1}, {0, 2, 2}, {1, 2, 5}, {1, 3, 1}, {2, 3, 8}})
}
