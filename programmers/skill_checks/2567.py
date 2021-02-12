def solution(scoville, K):
    answer = 0

    for i in range(len(scoville) - 1):
        if(min(scoville) >= K):
            break
        new = 0

        new += min(scoville)
        scoville.remove(new)
        
        new += min(scoville) * 2
        scoville.remove(min(scoville))
        
        scoville.append(new)
        answer += 1
        
    if min(scoville) < K:
        answer = -1
    # print(answer)
    return answer

solution([1, 1, 1, 1, 1], 100)