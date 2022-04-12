package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()
	n, _ := strconv.Atoi(str)

	queue := []int{}
	for ; n > 0; n-- {
		scanner.Scan()
		arr := strings.Split(scanner.Text(), " ")
		command := arr[0]

		switch command {
		case "push":
			num, _ := strconv.Atoi(arr[1])
			queue = append(queue, num)
		case "pop":
			if len(queue) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(queue[0])
				queue = queue[1:]
			}
		case "size":
			fmt.Println(len(queue))
		case "empty":
			if len(queue) == 0 {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		case "front":
			if len(queue) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(queue[0])
			}
		case "back":
			if len(queue) == 0 {
				fmt.Println(-1)
			} else {
				fmt.Println(queue[len(queue)-1])
			}
		}
	}
}
