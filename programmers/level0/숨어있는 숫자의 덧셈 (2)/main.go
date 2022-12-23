import (
	"regexp"
	"strconv"
)

func solution(my_string string) (sum int) {
	regex := regexp.MustCompile("[a-zA-Z]")
	for _, s := range regex.Split(my_string, -1) {
		n, err := strconv.Atoi(s)
		if err == nil {
			sum += n
		}
	}
	return sum
}
