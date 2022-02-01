def solution(x, n):
    answer = [x]
    for i in range(n - 1):
        answer.append(answer[i] + x)
    return answer


print(solution(2, 5))
print(solution(4, 3))
print(solution(-4, 2))
