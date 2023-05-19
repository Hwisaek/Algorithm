package main

func solution(number []int) (count int) {
	for i := 0; i < len(number)-2; i++ {
		for j := i + 1; j < len(number)-1; j++ {
			for k := j + 1; k < len(number); k++ {
				if number[i]+number[j]+number[k] == 0 {
					count++
				}
			}
		}
	}
	return
}
