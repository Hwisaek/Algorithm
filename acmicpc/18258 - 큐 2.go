package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := ""
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	n, _ := strconv.Atoi(str)

	queue := list.New()
	for ; n > 0; n-- {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		command := arr[0]

		switch command {
		case "push":
			num, _ := strconv.Atoi(arr[1])
			queue.PushBack(num)
		case "pop":
			if queue.Len() == 0 {
				result += "-1\n"
			} else {
				result += fmt.Sprintf("%d\n", queue.Front().Value)
				queue.Remove(queue.Front())
			}
		case "size":
			result += fmt.Sprintf("%d\n", queue.Len())
		case "empty":
			if queue.Len() == 0 {
				result += "1\n"
			} else {
				result += "0\n"
			}
		case "front":
			if queue.Len() == 0 {
				result += "-1\n"
				fmt.Println(-1)
			} else {
				result += fmt.Sprintf("%d\n", queue.Front().Value)
			}
		case "back":
			if queue.Len() == 0 {
				result += "-1\n"
			} else {
				result += fmt.Sprintf("%d\n", queue.Back().Value)
			}
		}
	}
	fmt.Println(result)
}
