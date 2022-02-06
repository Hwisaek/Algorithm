def solution(citations):  # 논문 별 인용 횟수
    answer = 0

    citations.sort(reverse=True)  # 역순 정렬

    for idx, c in enumerate(citations, start=1):
        if c < idx:
            break
        else:
            answer = idx

    return answer


print(solution([3, 0, 6, 1, 5]))
