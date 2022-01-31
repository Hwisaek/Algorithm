package main

import "fmt"

func main() {
	num, count := 0, 0

	_, _ = fmt.Scan(&num)
	origin := num
	for num = num%10*10 + (num/10+num%10)%10; ; num = num%10*10 + (num/10+num%10)%10 {
		count++
		if origin == num {
			break
		}
	}
	fmt.Println(count)
}
