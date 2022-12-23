func solution(myStr string, n int) (result []string) {
	for n < len(myStr) {
		result = append(result, myStr[:n])
		myStr = myStr[n:]
	}
	result = append(result, myStr)
	return result
}
