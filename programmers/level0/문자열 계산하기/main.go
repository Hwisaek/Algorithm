import (
	"strconv"
	"strings"
)

func solution(my_string string) (result int) {
	arr := strings.Split(my_string, " ")
	result, _ = strconv.Atoi(arr[0])
	for i := 1; i < len(arr); i += 2 {
		switch arr[i] {
		case "+":
			n, _ := strconv.Atoi(arr[i+1])
			result += n
		case "-":
			n, _ := strconv.Atoi(arr[i+1])
			result -= n
		}
	}
	return
}
