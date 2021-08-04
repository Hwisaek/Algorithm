def solution(n):
    answer, prev = 1, 0
    for i in range(n - 1):
        answer, prev = answer + prev, answer
    return answer % 1234567


# solution(4)
solution(3)
solution(100000)
