package 직사각형_별찍기

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for i := 0; i < b; i++ {
		row := ""
		for j := 0; j < a; j++ {
			row += "*"
		}
		fmt.Println(row)
	}
}
