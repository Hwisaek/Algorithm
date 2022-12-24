package main

func solution(coupon int) (count int) {
	for coupon >= 10 {
		count += coupon / 10
		coupon = coupon/10 + coupon%10
	}
	return
}

func main() {
	solution(100)
	solution(1081)
}
