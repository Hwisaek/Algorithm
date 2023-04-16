package main

func solution(cards1 []string, cards2 []string, goal []string) string {
	for _, v := range goal {
		switch {
		case len(cards1) > 0 && cards1[0] == v:
			cards1 = cards1[1:]
		case len(cards2) > 0 && cards2[0] == v:
			cards2 = cards2[1:]
		default:
			return "No"
		}
	}
	return "Yes"
}

func main() {
	solution([]string{"i", "drink", "water"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"})
	solution([]string{"i", "water", "drink"}, []string{"want", "to"}, []string{"i", "want", "to", "drink", "water"})
}
