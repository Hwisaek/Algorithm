package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, m, v int
	fmt.Fscan(rd, &n, &m, &v)
	graph := make([][]int, n+1)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscan(rd, &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	for _, intArr := range graph {
		sort.Ints(intArr)
	}
	_, _ = wr.WriteString(solution1260(graph, v))
	_ = wr.Flush()
}

func solution1260(graph [][]int, v int) (result string) {

	visited1 := make([]bool, len(graph))
	visited2 := make([]bool, len(graph))

	dfs(graph, v, visited1)
	fmt.Println()
	bfs(graph, v, visited2)
	return
}

func dfs(graph [][]int, n int, visited []bool) {
	if visited[n] {
		return
	}

	visited[n] = true
	fmt.Print(n, " ")

	if len(graph[n]) > 0 {
		for _, m := range graph[n] {
			if !visited[m] {
				dfs(graph, m, visited)
			}
		}
	}
}

func bfs(graph [][]int, n int, visited []bool) {
	queue := make([]int, 0)
	queue = append(queue, n)

	visited[n] = true

	for len(queue) > 0 {
		next := queue[0]
		queue = queue[1:]
		fmt.Print(next, " ")
		if len(graph[n]) > 0 {
			for _, i := range graph[next] {
				if !visited[i] {
					queue = append(queue, i)
					visited[i] = true
				}
			}
		}
	}
}
