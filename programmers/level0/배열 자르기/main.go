package 배열_자르기

func solution(numbers []int, num1 int, num2 int) []int {
	return numbers[num1 : num2+1]
}
