package main

func solution(players []string, callings []string) []string {
	idxMap := map[string]int{}
	for i, player := range players {
		idxMap[player] = i
	}

	for _, calling := range callings {
		callingIdx := idxMap[calling]
		forwarder := players[callingIdx-1]

		players[callingIdx], players[callingIdx-1] = players[callingIdx-1], players[callingIdx]
		idxMap[calling], idxMap[forwarder] = callingIdx-1, callingIdx
	}

	return players
}
