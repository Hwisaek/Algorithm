package main

import (
	"fmt"
	"strings"
)

var (
	n  int
	dx = []int{1, 0, -1, 0}
	dy = []int{0, 1, 0, -1}
)

func main() {
	var count1, count2 int
	var visited1, visited2 [][]bool

	fmt.Scanf("%d\n", &n)

	visited1 = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited1[i] = make([]bool, n)
	}
	visited2 = make([][]bool, n)
	for i := 0; i < n; i++ {
		visited2[i] = make([]bool, n)
	}

	graph := make([][]string, n)
	for i := 0; i < n; i++ {
		str := ""
		fmt.Scanf("%s\n", &str)
		graph[i] = strings.Split(str, "")
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if bfs(graph, visited1, i, j, false) {
				count1++
			}
			if bfs(graph, visited2, i, j, true) {
				count2++
			}
		}
	}

	fmt.Println(count1, count2)
}

func bfs(graph [][]string, visited [][]bool, x, y int, mode bool) (ok bool) {
	if visited[x][y] {
		return
	}

	visited[x][y] = true
	for i := 0; i < 4; i++ {
		x2 := x + dx[i]
		y2 := y + dy[i]

		if x2 < 0 || x2 >= n || y2 < 0 || y2 >= n {
			continue
		}

		now := graph[x][y]
		next := graph[x2][y2]

		switch mode {
		case true:
			if now == "G" {
				now = "R"
			}
			if next == "G" {
				next = "R"
			}
		}

		if now == next {
			bfs(graph, visited, x2, y2, mode)
		}
	}
	return true
}
