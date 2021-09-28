def solution(n):
    answer = n ** 0.5
    return (answer + 1) ** 2 if answer == int(answer) else -1
