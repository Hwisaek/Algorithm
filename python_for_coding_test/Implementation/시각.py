def solution(n):  # n: 시각
    # 0시부터 n시 59분 59초까지 3이 하나라도 포함되는 모든 경우의 수
    answer = 0

    for h in range(n + 1):
        for m in range(60):
            for s in range(60):
                if '3' in str(h) + str(m) + str(s):
                    answer += 1

    return answer


print(solution(5))
