import heapq


def solution(scoville, K):
    answer = 0
    heapq.heapify(scoville)

    while len(scoville) > 1:
        if scoville[0] >= K:
            return answer
        s = heapq.heappop(scoville) + heapq.heappop(scoville) * 2
        heapq.heappush(scoville, s)
        answer += 1
    return answer if scoville[0] >= K else -1


print(solution([1, 2, 3, 9, 10, 12], 7))
print(solution([1, 2, 3, 9, 10, 12], 9999999))
