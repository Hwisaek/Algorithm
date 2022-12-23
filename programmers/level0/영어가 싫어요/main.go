import "strconv"

func solution(numbers string) (result int) {
	words := map[string]byte{
		"zero":   0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	var answer []byte
	for i := 3; len(numbers) > 0; i++ {
		if n, ok := words[numbers[:i]]; ok {
			answer = append(answer, n+'0')
			numbers = numbers[i:]
			i = 2
		}
	}
	result, _ = strconv.Atoi(string(answer))
	return
}
