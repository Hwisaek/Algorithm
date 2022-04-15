package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	text, _ := rd.ReadString('\n') // 여기서 text는 마지막에 줄바꿈 문자를 포함하므로
	text = strings.TrimSpace(text) // 줄바꿈 문자를 제거해야 함
	n, _ := strconv.Atoi(text)

	queue := newQueue()
	for ; n > 0; n-- {
		str, _ := rd.ReadString('\n') // 여기서 str는 마지막에 줄바꿈 문자를 포함하므로
		str = strings.TrimSpace(str)  // 줄바꿈 문자를 제거해야 함
		arr := strings.Split(str, " ")
		command := arr[0]

		switch command {
		case "push":
			num, _ := strconv.Atoi(arr[1])
			queue.push(num)
		case "pop":
			wr.WriteString(strconv.Itoa(queue.pop()) + "\n")
		case "size":
			wr.WriteString(strconv.Itoa(queue.size()) + "\n")
		case "empty":
			wr.WriteString(strconv.Itoa(queue.empty()) + "\n")
		case "front":
			wr.WriteString(strconv.Itoa(queue.front()) + "\n")
		case "back":
			wr.WriteString(strconv.Itoa(queue.back()) + "\n")
		}
	}

}

type queue struct {
	queue []int
}

func newQueue() queue {
	return queue{[]int{}}
}

func (q *queue) push(x int) {
	q.queue = append(q.queue, x)
}

func (q *queue) pop() (n int) {
	if len(q.queue) == 0 {
		return -1
	}
	n = q.queue[0]
	q.queue = q.queue[1:]
	return
}

func (q *queue) size() int {
	return len(q.queue)
}

func (q *queue) empty() (n int) {
	if len(q.queue) == 0 {
		return 1
	}
	return 0
}

func (q *queue) front() int {
	if q.empty() == 0 {
		return q.queue[0]
	}
	return -1
}

func (q *queue) back() int {
	if q.empty() == 0 {
		return q.queue[q.size()-1]
	}
	return -1
}
