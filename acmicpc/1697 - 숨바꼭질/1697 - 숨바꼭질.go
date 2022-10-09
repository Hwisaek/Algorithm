package main

import (
	"bufio"
	. "fmt"
	"os"
)

var (
	r   = bufio.NewReader(os.Stdin)
	max = 100001
)

type point struct {
	x, count int
}

func main() {
	var n, k int
	Fscan(r, &n, &k)

	visited := make([]bool, max)

	q := []point{{
		x:     n,
		count: 0,
	}}

	for len(q) > 0 {
		now := q[0]
		q = q[1:]

		if visited[now.x] {
			continue
		}
		visited[now.x] = true

		if now.x == k {
			Println(now.count)
			break
		}

		if now.x-1 >= 0 {
			p := point{
				x:     now.x - 1,
				count: now.count + 1,
			}
			q = append(q, p)
		}
		if now.x+1 < max {
			p := point{
				x:     now.x + 1,
				count: now.count + 1,
			}
			q = append(q, p)
		}
		if now.x*2 < max {
			p := point{
				x:     now.x * 2,
				count: now.count + 1,
			}
			q = append(q, p)
		}
	}
}
