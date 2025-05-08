package solution

func Solution(N int) (answer int) {
	for ; N%2 == 0; answer++ {
		N /= 2
	}
	
	return
}
