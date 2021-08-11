def solution(n, fears):
    answer = 0
    fears.sort(reverse=True)  # 공포도를 역 순으로 정렬

    while fears:  # 공포도가 존재하면 반복
        now = 0  # 현재 그룹원 인원 수
        while True:  # 그룹 생성
            limit = fears.pop()  # 해당 그룹의 공포도 최댓값
            now += 1  # 현재 그룹원 +1
            if now >= limit:  # 그룹에 인원이 꽉차면 출발
                break
            if len(fears) < limit - now:  # 남은 인원이 부족하면 마을 잔류
                return answer
        answer += 1

    return answer


print(solution(5, [2, 3, 1, 2, 2]))
print(solution(5, [2, 1, 2, 4, 5]))
print(solution(5, [1, 1, 1, 1, 1]))
