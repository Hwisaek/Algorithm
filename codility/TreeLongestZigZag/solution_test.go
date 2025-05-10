package solution

type Tree struct {
	X int
	L *Tree
	R *Tree
}

var answer int = 0

// 노드 수 N: 1 ~ 10만
// 최대 트리 높이: 800
func Solution(T *Tree) int {
	dfs(T, 0, 0)
	return answer
}

func dfs(node *Tree, dir int, turn int) {
	if node == nil {
		return
	}

	lturn := turn
	rturn := turn

	switch dir {
	case -1:
		rturn++
	case 1:
		lturn++
	}

	dfs(node.L, -1, lturn)
	dfs(node.R, 1, rturn)

	answer = max(answer, turn)
}
