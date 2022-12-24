package main

func solution(spell []string, dic []string) int {
	letters := make(map[string]int)
	for _, letter := range spell {
		letters[letter]++
	}

	for _, word := range dic {
		wordLetters := make(map[string]int)
		for _, letter := range word {
			wordLetters[string(letter)]++
		}

		isValid := true
		for letter, count := range letters {
			if wordLetters[letter] != count {
				isValid = false
				break
			}
		}

		if isValid {
			return 1
		}
	}

	return 2
}

func main() {
	//solution([]string{"p", "o", "s"}, []string{"sod", "eocd", "qixm", "adio", "soo"})
	//solution([]string{"z", "d", "x"}, []string{"def", "dww", "dzx", "loveaw"})
	solution([]string{"s", "o", "m", "d"}, []string{"moos", "dzx", "smm", "sunmmo", "som"})
}
