package 직사각형_별찍기

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result := ""
	for i := 0; i < b; i++ {
		for j := 0; j < a; j++ {
			result += "*"
		}
		result += "\n"
	}
	fmt.Println(result)
}
