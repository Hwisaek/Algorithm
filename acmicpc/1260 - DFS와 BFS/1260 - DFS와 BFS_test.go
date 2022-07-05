package _260___DFS와_BFS

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)

	var n, m, v int
	fmt.Fscanf(rd, "%d %d %d", &n, &m, &v)

	graph := make([][]int, n)
	for i := 0; i < m; i++ {
		var a, b int
		fmt.Fscanf(rd, "%d %d", &a, &b)
		graph[a] = append(graph[a], b)
		graph[b] = append(graph[b], a)
	}

	_, _ = wr.WriteString(solution1260(graph, v))
	_ = wr.Flush()
}

func solution1260(graph [][]int, v int) (result string) {
	result += dfs(graph, v, false) + "\n" + dfs(graph, v, true)
	return
}

func dfs(graph [][]int, v int, isBfs bool) (result string) {
	visited := make([]bool, len(graph))
	visited[v] = true

	queue := make([]int, 0)
	queue = append(queue, v)

	for len(queue) > 0 {
		if isBfs {
			v = queue[0]
			queue = queue[1:]
		} else {
			v = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		}
		result += fmt.Sprintf("%d ", v)
		for _, w := range graph[v] {
			if !visited[w] {
				visited[w] = true
				queue = append(queue, w)
			}
		}
	}
	return
}
