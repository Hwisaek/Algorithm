package 옷가게_할인_받기

func solution(price int) int {
	switch {
	case price >= 500000:
		return price * 80 / 100
	case price >= 300000:
		return price * 90 / 100
	case price >= 100000:
		return price * 95 / 100
	default:
		return price
	}
}
