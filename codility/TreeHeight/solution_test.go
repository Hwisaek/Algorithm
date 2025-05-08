package solution

import (
	"testing"
)

type Tree struct {
	X int
	L *Tree
	R *Tree
}

func Solution(T *Tree) (answer int) {
	if T.L != nil {
		answer = max(answer, Solution(T.L)+1)
	}
	if T.R != nil {
		answer = max(answer, Solution(T.R)+1)
	}

	return answer
}

func TestSolution(t *testing.T) {
	tree2 := &Tree{
		X: 5,
		L: &Tree{
			X: 3,
			L: &Tree{
				X: 20,
			},
			R: &Tree{
				X: 21,
			},
		},
		R: &Tree{
			X: 10,
			L: &Tree{
				X: 1,
			},
		},
	}

	if Solution(tree2) != 3 {
		t.Errorf("got %d", Solution(tree2))
	}
}
