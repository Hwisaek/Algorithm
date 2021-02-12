def solution(citations):
    answer = 0
    citations.sort(reverse=True)
    for i in range(len(citations)):
        for j in range(i + 1):
            if citations[j] < i + 1:
                break
        else:
            answer = i + 1
    return answer

print(solution([1, 2]))